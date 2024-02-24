package constants

const (
	LecturerStatusAktif               = "AKTIF"
	LecturerStatusCuti                = "CUTI"
	LecturerStatusIjinBelajar         = "IJIN BELAJAR"
	LecturerStatusKeluar              = "KELUAR"
	LecturerStatusAlmarhum            = "ALMARHUM"
	LecturerStatusTidakAktif          = "TIDAK AKTIF"
	LecturerStatusPensiun             = "PENSIUN"
	LecturerStatusTugasBelajar        = "TUGAS BELAJAR"
	LecturerStatusTugasDiInstansiLain = "TUGAS DI INSTANSI LAIN"
	LecturerStatusLainnya             = "LAINNYA"
	LecturerStatusGantiNidn           = "GANTI NIDN"
	LecturerStatusHapusNidn           = "HAPUS NIDN"
)

func LecturerResignStatus() []string {
	return []string{
		LecturerStatusKeluar,
		LecturerStatusAlmarhum,
		LecturerStatusPensiun,
	}
}
