package router

import (
	"example.com/rest-api-recoll-mobile/controller"
	"example.com/rest-api-recoll-mobile/helper"
	"example.com/rest-api-recoll-mobile/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery(), helper.Logger())

	// Middleware to set X-Content-Type-Options header
	router.Use(func(c *gin.Context) {
		// Set X-Content-Type-Options header
		c.Header("X-Content-Type-Options", "nosniff")

		c.Header("Cache-Control", "no-store")
		c.Header("Pragma", "no-cache")

		c.Next()
	})

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*", "http://localhost:8081", "https://6yd68t.csb.app"},
	// 	AllowHeaders:     []string{"X-Requested-with, Content-Type, Authorization, Access-Control-Allow-Origin"},
	// 	AllowMethods:     []string{"POST, OPTIONS, GET, PUT, DELETE, OPTIONS"},
	// 	AllowCredentials: true,
	// }))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := router.Group("/api")
	{

		public := api.Group("/public")
		{
			public.POST("/login", controller.Login)
			// public.GET("user/:id", controllers.GetUserById)
			// / public.POST("/signup", controllers.Signup)
		}

		version := api.Group("/version")
		{
			version.GET("/cek/:version", controller.VersionCheck)
		}

		global := api.Group("/global").Use(middleware.Authz())
		{
			global.GET("/cabang/list/:kode", controller.ListCabang)
			global.GET("/unit/list/:kode", controller.ListUnitByCabang)
			global.GET("/wilayah", controller.ListWilayah)

			global.GET("/cara-penanganan/npl/list", controller.ListCaraPenanganan)
			global.GET("/cara-penanganan/wo/list", controller.ListCaraPenangananWo)
			global.GET("/rencana-penanganan/npl/list/:kode", controller.ListRencanaPenyelesaian)
			global.GET("/rencana-penanganan/wo/list/:kode", controller.ListRencanaPenyelesaianWo)
			global.GET("/pic-penanganan/list/:kode", controller.ListPicPenanganan)
			global.GET("/pic-penanganan/detail/:kode", controller.DetailPic)
			global.GET("/cabang/approval/list", controller.ListCabangApproval)
			global.GET("/penyelesaian/list", controller.ListPenyelesaian)
			global.GET("/hasil-penyelesaian/list/:kode", controller.ListHasilPenyelesaian)
		}
		profile := api.Group("/profile").Use(middleware.Authz())
		{
			profile.GET("/profile/:id", controller.Profile)
		}

		deskCall := api.Group("/desk-call").Use(middleware.Authz())
		{
			// ROUTER DESK CALL
			deskCall.GET("/list/kode/:kode", controller.ListDeskCall)
			deskCall.GET("/detail", controller.DetailDeskCall)
			deskCall.GET("/riwayat", controller.RiwayatDeskCall)
			deskCall.GET("/list-nasabah/kode/:kode", controller.ListNasabahDeskCall)
			deskCall.POST("/insert", controller.InsertDeskCall)
		}

		dca := api.Group("/dca").Use(middleware.Authz())
		{
			// ROUTER DCA
			dca.GET("/list/kode/:kode/type/:type/search/:search/page/:page/row/:row/sort/:sort", controller.ListMonitoringDCA)
			dca.GET("/lmpd/list/kode/:kode", controller.ListLMPD)
			dca.GET("/mppd/list/kode/:kode", controller.ListMPPD)
			dca.POST("/mppd/insert", controller.InsertMPPD)
		}
		master := api.Group("/master").Use(middleware.Authz())
		{
			// ROUTER DESK CALL
			master.GET("/nasabah/list/kode/:kode", controller.ListNasabah)
		}

		dashboard := api.Group("/dashboard").Use(middleware.Authz())
		{
			// ROUTER DESK CALL
			dashboard.GET("/penurunan-npl/list", controller.ListPenurunanNFL)
			dashboard.GET("/grand-total", controller.GrandTotal)
		}
		papeline := api.Group("/papeline").Use(middleware.Authz())
		{
			papeline.GET("/bucket/bulan/list", controller.ListTableBulan)

			// RMD PAPELINE - HARI - RESUME PLAN
			papeline.GET("/resume-plan/col/list", controller.ListResumePlanCol)
			papeline.GET("/resume-plan/wo/list", controller.ListResumePlanWo)

			// CB PAPELINE - HARI - RESUME PLAN
			papeline.GET("/resume-plan/cb/col/list", controller.ListCbResumePlanCol)
			papeline.GET("/resume-plan/cb/wo/list", controller.ListCbResumePlanWo)
			papeline.GET("/resume-plan/cb/status", controller.StatusPlan)

			// RMD PAPELINE - HARI - DETAIL PLAN
			papeline.GET("/detail-plan/npl/list", controller.ListDetailPlanNpl)
			papeline.GET("/detail-plan/wo/list", controller.ListDetailPlanWo)

			// CB PAPELINE - HARI - DETAIL PLAN
			papeline.GET("/detail-plan/cb/npl/list", controller.ListCbDetailPlanNpl)
			papeline.GET("/detail-plan/cb/npl/detail", controller.DetailPlanHariCB)
			papeline.GET("/detail-plan/cb/wo/list", controller.ListCbDetailPlanWo)
			papeline.GET("/detail-plan/cb/wo/detail", controller.DetailPlanWoHariCB)
			papeline.GET("/detail-plan/cb/nasabah/list", controller.ListNasabahPlanHari)
			papeline.GET("/detail-plan/cb/nasabah/detail", controller.DetailNasabahPlanHari)
			papeline.GET("/detail-plan/cb/nasabah/wo/list", controller.ListNasabahWoPlanHari)
			papeline.GET("/detail-plan/cb/nasabah/wo/detail", controller.DetailNasabahWoPlanHari)

			papeline.POST("/plan/hari/nasabah/submit", controller.SubmitPlanHari)
			papeline.POST("/plan/hari/kirim", controller.KirimPlanHari)
			papeline.POST("/plan/hari/penyelesaian/submit", controller.SubmitPeneyelesaian)

			// RMD PAPELINE - HARI - RESUME REALISASI
			papeline.GET("/resume-realisasi/col/list", controller.ListResumeRealisasiCol)
			papeline.GET("/resume-realisasi/wo/list", controller.ListResumeRealisasiWo)

			// CB PAPELINE - HARI - RESUME REALISASI
			papeline.GET("/resume-realisasi/cb/col/list", controller.ListCbResumeRealisasiCol)
			papeline.GET("/resume-realisasi/cb/wo/list", controller.ListCbResumeRealisasiWo)

			// RMD PAPELINE - HARI - DETAIL REALISASI
			papeline.GET("/detail-realisasi/npl/list", controller.ListDetailRealisasiNpl)
			papeline.GET("/detail-realisasi/wo/list", controller.ListDetailRealisasiWo)
			papeline.GET("/detail-realisasi/npl/detail", controller.DetailRealisasiNplById)
			papeline.GET("/detail-realisasi/wo/detail", controller.DetailRealisasiWoById)
			papeline.GET("/plan-hari/history", controller.ListHistoriPlanHari)

			// CB PAPELINE - HARI - DETAIL REALISASI
			papeline.GET("/detail-realisasi/cb/npl/list", controller.ListCbDetailRealisasiNpl)
			papeline.GET("/detail-realisasi/cb/wo/list", controller.ListCbDetailRealisasiWo)

			// RMD PAPELINE - BULAN -RESUME
			papeline.GET("/bulan/resume/list", controller.ListBulanResume)
			papeline.GET("/bulan/resume/npl/list", controller.ListBulanResumeNpl)
			papeline.GET("/bulan/resume/total/list", controller.ListBulanResumeTotal)
			papeline.GET("/bulan/resume/wo/list", controller.ListBulanResumeWo)

			// RMD PAPELINE - BULAN - DETAIL PLAN
			papeline.GET("/bulan/detail-plan/npl/list", controller.ListBulanDetailPlanNpl)
			papeline.GET("/bulan/detail-plan/npl/detail", controller.ListBulanDetailPlanNplById)
			papeline.GET("/bulan/detail-plan/wo/list", controller.ListBulanDetailPlanWo)
			papeline.GET("/bulan/detail-plan/wo/detail", controller.ListBulanDetailPlanWoById)

			// RMD PAPELINE - BULAN - CEK PLAN
			papeline.GET("/bulan/cek-plan/list", controller.ListBulanCekPlan)

			// RMD PAPELINE  - RECOVERY
			papeline.GET("/recovery/kolektibilitas/list", controller.ListKolektibilitasRecovery)
			papeline.GET("/recovery/wo/list", controller.ListWoRecovery)

			// CB PAPELINE - BULAN
			papeline.GET("/bulan/cb/kol3/list", controller.CBListPipelineBulanKol3)
			papeline.GET("/bulan/cb/kol4/list", controller.CBListPipelineBulanKol4)
			papeline.GET("/bulan/cb/kol5/list", controller.CBListPipelineBulanKol5)
			papeline.GET("/bulan/cb/wo/list", controller.CBListPipelineBulanWo)

			papeline.GET("/bulan/cb/resume/total", controller.CBTotalResume)
			papeline.GET("/bulan/cb/resume/table", controller.CBTableResume)

		}

		realisasi := api.Group("/realisasi").Use(middleware.Authz())
		{
			// RMD REALISASI - TRANSAKSI - RESUME
			realisasi.GET("/transaksi/resume/list", controller.ResumeTransaksiWidget)
			realisasi.GET("/transaksi/resume/table/list", controller.ListResumeTransaksi)

			// RMD REALISASI - TRANSAKSI - DETAIL TRANSAKSI
			realisasi.GET("/transaksi/detail-transaksi/list", controller.ListDetailTransaksi)
			realisasi.GET("/transaksi/detail-transaksi/detail", controller.DetailTransaksiById)

			// CAB REALISASI - TRANSAKSI - RESUME
			realisasi.GET("/transaksi/resume/cab/list", controller.ResumeTransaksiWidgetCab)
			realisasi.GET("/transaksi/resume/table/cab/list", controller.ListResumeTransaksiCab)

		}

		planbulan := api.Group("/planbulan")
		{
			planbulan.GET("/listkol", controller.GetPlanBulan)        //listplanbulan ?->query param
			planbulan.GET("/detail/:id", controller.GetPlanBulanById) //detailplanbulan path param
			planbulan.GET("/wo", controller.GetRecoveryWoById)        //list wo
			planbulan.PUT(":id", controller.UpdatePlanBulan)
			planbulan.POST("/saveplanbulan", controller.CreatePlanBulan)     //save detail planbulan
			planbulan.POST("/saveplanbulanwo", controller.CreatePlanBulanWo) //save detail planbulanwo
		}

		penanganan := api.Group("/penanganan")
		{
			penanganan.GET("", controller.ListPicPenangananCab)
			penanganan.GET(":id", controller.GetPicPenangananId)
		}

		resume := api.Group("/resume")
		{
			resume.GET("/baseonkol", controller.ResumeBaseonKol)
			resume.GET("/resumeplankol-1", controller.ResumePlanKol1)
			resume.GET("/resumeplankol-2", controller.ResumePlanKol2)
			resume.GET("/resumeplankol-3", controller.ResumePlanKol3)
			resume.GET("/resumeplankol-4", controller.ResumePlanKol4)
			resume.GET("/resumeplankol-5", controller.ResumePlanKol5)
			resume.GET("/baseonwo", controller.ResumeBaseonWo)
			resume.GET("/resumewo", controller.ResumePlanWo)
			resume.GET("/resumerembesnpl", controller.ResumeRembesNpl)
		}

	}

	return router
}
