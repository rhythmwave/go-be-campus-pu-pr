package constants

import (
	"context"
	"net/http"
)

const (
	PermissionAktivitasMahasiswa = "Aktivitas Mahasiswa"
	PermissionDataReferensi      = "Data Referensi"
	PermissionDosen              = "Dosen"
	PermissionEvaluasi           = "Evaluasi"
	PermissionHasilStudi         = "Hasil Studi"
	PermissionInputNilai         = "Input Nilai"
	PermissionJadwal             = "Jadwal"
	PermissionJatahSks           = "Jatah Sks"
	PermissionJenisNilai         = "Jenis Nilai"
	PermissionKelulusan          = "Kelulusan"
	PermissionKemahasiswaan      = "Kemahasiswaan"
	PermissionKomponenNilai      = "Komponen Nilai"
	PermissionKuisioner          = "Kuisioner"
	PermissionKurikulum          = "Kurikulum"
	PermissionLaporan            = "Laporan"
	PermissionManajemenRuang     = "Manajemen Ruang"
	PermissionMenungguUKOM       = "Menunggu UKOM"
	PermissionNilai              = "Nilai"
	PermissionPejabat            = "Pejabat"
	PermissionPenyelenggaraMBKM  = "Penyelenggara MBKM"
	PermissionPersetujuanKRS     = "Persetujuan KRS"
	PermissionPesertaKelas       = "Peserta Kelas"
	PermissionPesertaMBKM        = "Peserta MBKM"
	PermissionPortalAkademik     = "Portal Akademik"
	PermissionPredikatKelulusan  = "Predikat Kelulusan"
	PermissionProgramStudi       = "Program Studi"
	PermissionRekapNilai         = "Rekap Nilai"
	PermissionRencanaStudi       = "Rencana Studi"
	PermissionSKPI               = "SKPI"
	PermissionSemester           = "Semester"
	PermissionStatusRegistrasi   = "Status Registrasi"
	PermissionTranskrip          = "Transkrip"
	PermissionTugasAkhir         = "Tugas Akhir"
	PermissionVirtualClass       = "Virtual Class"
)

func PermissionAktivitasMahasiswaMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionAktivitasMahasiswa)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionDataReferensiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionDataReferensi)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionDosenMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionDosen)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionEvaluasiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionEvaluasi)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionHasilStudiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionHasilStudi)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionInputNilaiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionInputNilai)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionJadwalMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionJadwal)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionJatahSksMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionJatahSks)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionJenisNilaiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionJenisNilai)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionKelulusanMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionKelulusan)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionKemahasiswaanMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionKemahasiswaan)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionKomponenNilaiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionKomponenNilai)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionKuisionerMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionKuisioner)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionKurikulumMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionKurikulum)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionLaporanMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionLaporan)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionManajemenRuangMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionManajemenRuang)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionMenungguUKOMMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionMenungguUKOM)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionNilaiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionNilai)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionPejabatMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionPejabat)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionPenyelenggaraMBKMMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionPenyelenggaraMBKM)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionPersetujuanKRSMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionPersetujuanKRS)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionPesertaKelasMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionPesertaKelas)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionPesertaMBKMMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionPesertaMBKM)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionPortalAkademikMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionPortalAkademik)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionPredikatKelulusanMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionPredikatKelulusan)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionProgramStudiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionProgramStudi)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionRekapNilaiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionRekapNilai)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionRencanaStudiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionRencanaStudi)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionSKPIMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionSKPI)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionSemesterMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionSemester)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionStatusRegistrasiMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionStatusRegistrasi)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionTranskripMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionTranskrip)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionTugasAkhirMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionTugasAkhir)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func PermissionVirtualClassMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, PermissionContextKey, PermissionVirtualClass)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}
