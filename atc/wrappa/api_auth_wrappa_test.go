package wrappa_test

import (
	"github.com/concourse/concourse/atc/types"
	"net/http"

	"github.com/concourse/concourse/atc/api/auth"
	"github.com/concourse/concourse/atc/db/dbfakes"
	"github.com/concourse/concourse/atc/wrappa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/rata"
)

var _ = Describe("APIAuthWrappa", func() {
	var (
		rejector                                auth.Rejector
		fakeCheckPipelineAccessHandlerFactory   auth.CheckPipelineAccessHandlerFactory
		fakeCheckBuildReadAccessHandlerFactory  auth.CheckBuildReadAccessHandlerFactory
		fakeCheckBuildWriteAccessHandlerFactory auth.CheckBuildWriteAccessHandlerFactory
		fakeCheckWorkerTeamAccessHandlerFactory auth.CheckWorkerTeamAccessHandlerFactory
		fakeBuildFactory                        *dbfakes.FakeBuildFactory
	)

	BeforeEach(func() {
		fakeTeamFactory := new(dbfakes.FakeTeamFactory)
		workerFactory := new(dbfakes.FakeWorkerFactory)
		fakeBuildFactory = new(dbfakes.FakeBuildFactory)
		fakeCheckPipelineAccessHandlerFactory = auth.NewCheckPipelineAccessHandlerFactory(
			fakeTeamFactory,
		)
		rejector = auth.UnauthorizedRejector{}

		fakeCheckBuildReadAccessHandlerFactory = auth.NewCheckBuildReadAccessHandlerFactory(fakeBuildFactory)
		fakeCheckBuildWriteAccessHandlerFactory = auth.NewCheckBuildWriteAccessHandlerFactory(fakeBuildFactory)
		fakeCheckWorkerTeamAccessHandlerFactory = auth.NewCheckWorkerTeamAccessHandlerFactory(workerFactory)
	})

	authenticateIfTokenProvided := func(handler http.Handler) http.Handler {
		return auth.CheckAuthenticationIfProvidedHandler(
			handler,
			rejector,
		)
	}

	authenticated := func(handler http.Handler) http.Handler {
		return auth.CheckAuthenticationHandler(
			handler,
			rejector,
		)
	}

	authenticatedAndAdmin := func(handler http.Handler) http.Handler {
		return auth.CheckAdminHandler(
			handler,
			rejector,
		)
	}

	authorized := func(handler http.Handler) http.Handler {
		return auth.CheckAuthorizationHandler(
			handler,
			rejector,
		)
	}

	openForPublicPipelineOrAuthorized := func(handler http.Handler) http.Handler {
		return fakeCheckPipelineAccessHandlerFactory.HandlerFor(
			handler,
			rejector,
		)
	}

	doesNotCheckIfPrivateJob := func(handler http.Handler) http.Handler {
		return fakeCheckBuildReadAccessHandlerFactory.AnyJobHandler(
			handler,
			rejector,
		)
	}

	checksIfPrivateJob := func(handler http.Handler) http.Handler {
		return fakeCheckBuildReadAccessHandlerFactory.CheckIfPrivateJobHandler(
			handler,
			rejector,
		)
	}

	checkWritePermissionForBuild := func(handler http.Handler) http.Handler {
		return fakeCheckBuildWriteAccessHandlerFactory.HandlerFor(
			handler,
			rejector,
		)
	}

	checkTeamAccessForWorker := func(handler http.Handler) http.Handler {
		return fakeCheckWorkerTeamAccessHandlerFactory.HandlerFor(
			handler,
			rejector,
		)
	}

	Describe("Wrap", func() {
		var (
			inputHandlers    rata.Handlers
			expectedHandlers rata.Handlers

			wrappedHandlers rata.Handlers
		)

		BeforeEach(func() {
			inputHandlers = rata.Handlers{}

			for _, route := range types.Routes {
				inputHandlers[route.Name] = &stupidHandler{}
			}

			expectedHandlers = rata.Handlers{

				// authorized or public pipeline
				types.GetBuild:       doesNotCheckIfPrivateJob(inputHandlers[types.GetBuild]),
				types.BuildResources: doesNotCheckIfPrivateJob(inputHandlers[types.BuildResources]),

				// authorized or public pipeline and public job
				types.BuildEvents:         checksIfPrivateJob(inputHandlers[types.BuildEvents]),
				types.ListBuildArtifacts:  checksIfPrivateJob(inputHandlers[types.ListBuildArtifacts]),
				types.GetBuildPreparation: checksIfPrivateJob(inputHandlers[types.GetBuildPreparation]),
				types.GetBuildPlan:        checksIfPrivateJob(inputHandlers[types.GetBuildPlan]),

				// resource belongs to authorized team
				types.AbortBuild: checkWritePermissionForBuild(inputHandlers[types.AbortBuild]),

				// resource belongs to authorized team
				types.PruneWorker:              checkTeamAccessForWorker(inputHandlers[types.PruneWorker]),
				types.LandWorker:               checkTeamAccessForWorker(inputHandlers[types.LandWorker]),
				types.ReportWorkerContainers:   checkTeamAccessForWorker(inputHandlers[types.ReportWorkerContainers]),
				types.ReportWorkerVolumes:      checkTeamAccessForWorker(inputHandlers[types.ReportWorkerVolumes]),
				types.RetireWorker:             checkTeamAccessForWorker(inputHandlers[types.RetireWorker]),
				types.ListDestroyingContainers: checkTeamAccessForWorker(inputHandlers[types.ListDestroyingContainers]),
				types.ListDestroyingVolumes:    checkTeamAccessForWorker(inputHandlers[types.ListDestroyingVolumes]),

				// belongs to public pipeline or authorized
				types.GetPipeline:                   openForPublicPipelineOrAuthorized(inputHandlers[types.GetPipeline]),
				types.GetJobBuild:                   openForPublicPipelineOrAuthorized(inputHandlers[types.GetJobBuild]),
				types.PipelineBadge:                 openForPublicPipelineOrAuthorized(inputHandlers[types.PipelineBadge]),
				types.JobBadge:                      openForPublicPipelineOrAuthorized(inputHandlers[types.JobBadge]),
				types.ListJobs:                      openForPublicPipelineOrAuthorized(inputHandlers[types.ListJobs]),
				types.GetJob:                        openForPublicPipelineOrAuthorized(inputHandlers[types.GetJob]),
				types.ListJobBuilds:                 openForPublicPipelineOrAuthorized(inputHandlers[types.ListJobBuilds]),
				types.ListPipelineBuilds:            openForPublicPipelineOrAuthorized(inputHandlers[types.ListPipelineBuilds]),
				types.GetResource:                   openForPublicPipelineOrAuthorized(inputHandlers[types.GetResource]),
				types.ListBuildsWithVersionAsInput:  openForPublicPipelineOrAuthorized(inputHandlers[types.ListBuildsWithVersionAsInput]),
				types.ListBuildsWithVersionAsOutput: openForPublicPipelineOrAuthorized(inputHandlers[types.ListBuildsWithVersionAsOutput]),
				types.ListResources:                 openForPublicPipelineOrAuthorized(inputHandlers[types.ListResources]),
				types.ListResourceTypes:             openForPublicPipelineOrAuthorized(inputHandlers[types.ListResourceTypes]),
				types.ListResourceVersions:          openForPublicPipelineOrAuthorized(inputHandlers[types.ListResourceVersions]),
				types.GetResourceCausality:          openForPublicPipelineOrAuthorized(inputHandlers[types.GetResourceCausality]),
				types.GetResourceVersion:            openForPublicPipelineOrAuthorized(inputHandlers[types.GetResourceVersion]),

				// authenticated
				types.CreateBuild:     authenticated(inputHandlers[types.CreateBuild]),
				types.GetContainer:    authenticated(inputHandlers[types.GetContainer]),
				types.HijackContainer: authenticated(inputHandlers[types.HijackContainer]),
				types.ListContainers:  authenticated(inputHandlers[types.ListContainers]),
				types.ListVolumes:     authenticated(inputHandlers[types.ListVolumes]),
				types.ListTeamBuilds:  authenticated(inputHandlers[types.ListTeamBuilds]),
				types.ListWorkers:     authenticated(inputHandlers[types.ListWorkers]),
				types.RegisterWorker:  authenticated(inputHandlers[types.RegisterWorker]),
				types.HeartbeatWorker: authenticated(inputHandlers[types.HeartbeatWorker]),
				types.DeleteWorker:    authenticated(inputHandlers[types.DeleteWorker]),
				types.GetTeam:         authenticated(inputHandlers[types.GetTeam]),
				types.SetTeam:         authenticated(inputHandlers[types.SetTeam]),
				types.RenameTeam:      authenticated(inputHandlers[types.RenameTeam]),
				types.DestroyTeam:     authenticated(inputHandlers[types.DestroyTeam]),
				types.GetUser:         authenticated(inputHandlers[types.GetUser]),

				//authenticateIfTokenProvided / delegating to handler
				types.GetInfo:              authenticateIfTokenProvided(inputHandlers[types.GetInfo]),
				types.GetCheck:             authenticateIfTokenProvided(inputHandlers[types.GetCheck]),
				types.DownloadCLI:          authenticateIfTokenProvided(inputHandlers[types.DownloadCLI]),
				types.CheckResourceWebHook: authenticateIfTokenProvided(inputHandlers[types.CheckResourceWebHook]),
				types.ListAllPipelines:     authenticateIfTokenProvided(inputHandlers[types.ListAllPipelines]),
				types.ListBuilds:           authenticateIfTokenProvided(inputHandlers[types.ListBuilds]),
				types.ListPipelines:        authenticateIfTokenProvided(inputHandlers[types.ListPipelines]),
				types.ListAllJobs:          authenticateIfTokenProvided(inputHandlers[types.ListAllJobs]),
				types.ListAllResources:     authenticateIfTokenProvided(inputHandlers[types.ListAllResources]),
				types.ListTeams:            authenticateIfTokenProvided(inputHandlers[types.ListTeams]),
				types.MainJobBadge:         authenticateIfTokenProvided(inputHandlers[types.MainJobBadge]),
				types.GetWall:              authenticateIfTokenProvided(inputHandlers[types.GetWall]),

				// authenticated and is admin
				types.GetLogLevel:          authenticatedAndAdmin(inputHandlers[types.GetLogLevel]),
				types.SetLogLevel:          authenticatedAndAdmin(inputHandlers[types.SetLogLevel]),
				types.GetInfoCreds:         authenticatedAndAdmin(inputHandlers[types.GetInfoCreds]),
				types.ListActiveUsersSince: authenticatedAndAdmin(inputHandlers[types.ListActiveUsersSince]),
				types.SetWall:              authenticatedAndAdmin(inputHandlers[types.SetWall]),
				types.ClearWall:            authenticatedAndAdmin(inputHandlers[types.ClearWall]),

				// authorized (requested team matches resource team)
				types.CheckResource:           authorized(inputHandlers[types.CheckResource]),
				types.CheckResourceType:       authorized(inputHandlers[types.CheckResourceType]),
				types.CreateJobBuild:          authorized(inputHandlers[types.CreateJobBuild]),
				types.RerunJobBuild:           authorized(inputHandlers[types.RerunJobBuild]),
				types.DeletePipeline:          authorized(inputHandlers[types.DeletePipeline]),
				types.DisableResourceVersion:  authorized(inputHandlers[types.DisableResourceVersion]),
				types.EnableResourceVersion:   authorized(inputHandlers[types.EnableResourceVersion]),
				types.PinResourceVersion:      authorized(inputHandlers[types.PinResourceVersion]),
				types.UnpinResource:           authorized(inputHandlers[types.UnpinResource]),
				types.SetPinCommentOnResource: authorized(inputHandlers[types.SetPinCommentOnResource]),
				types.GetConfig:               authorized(inputHandlers[types.GetConfig]),
				types.GetCC:                   authorized(inputHandlers[types.GetCC]),
				types.GetVersionsDB:           authorized(inputHandlers[types.GetVersionsDB]),
				types.ListJobInputs:           authorized(inputHandlers[types.ListJobInputs]),
				types.OrderPipelines:          authorized(inputHandlers[types.OrderPipelines]),
				types.PauseJob:                authorized(inputHandlers[types.PauseJob]),
				types.PausePipeline:           authorized(inputHandlers[types.PausePipeline]),
				types.ArchivePipeline:         authorized(inputHandlers[types.ArchivePipeline]),
				types.RenamePipeline:          authorized(inputHandlers[types.RenamePipeline]),
				types.SaveConfig:              authorized(inputHandlers[types.SaveConfig]),
				types.UnpauseJob:              authorized(inputHandlers[types.UnpauseJob]),
				types.ScheduleJob:             authorized(inputHandlers[types.ScheduleJob]),
				types.UnpausePipeline:         authorized(inputHandlers[types.UnpausePipeline]),
				types.ExposePipeline:          authorized(inputHandlers[types.ExposePipeline]),
				types.HidePipeline:            authorized(inputHandlers[types.HidePipeline]),
				types.CreatePipelineBuild:     authorized(inputHandlers[types.CreatePipelineBuild]),
				types.ClearTaskCache:          authorized(inputHandlers[types.ClearTaskCache]),
				types.CreateArtifact:          authorized(inputHandlers[types.CreateArtifact]),
				types.GetArtifact:             authorized(inputHandlers[types.GetArtifact]),
			}
		})

		JustBeforeEach(func() {
			wrappedHandlers = wrappa.NewAPIAuthWrappa(
				fakeCheckPipelineAccessHandlerFactory,
				fakeCheckBuildReadAccessHandlerFactory,
				fakeCheckBuildWriteAccessHandlerFactory,
				fakeCheckWorkerTeamAccessHandlerFactory,
			).Wrap(inputHandlers)

		})

		It("validates sensitive routes, and noop validates public routes", func() {
			for name, _ := range inputHandlers {
				Expect(wrappedHandlers[name]).To(BeIdenticalTo(expectedHandlers[name]))
			}
		})
	})
})
