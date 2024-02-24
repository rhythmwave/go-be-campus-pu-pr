package constants

const (
	StudentStatusAktif        = "AKTIF"
	StudentStatusCuti         = "CUTI"
	StudentStatusKeluar       = "KELUAR"
	StudentStatusLulus        = "LULUS"
	StudentStatusNonAktif     = "NON-AKTIF"
	StudentStatusDropOut      = "DROP-OUT"
	StudentStatusMenungguUkom = "MENUNGGU-UKOM"
	StudentStatusMbkm         = "MBKM"
)

func PersistentStudentStatus() []string {
	return []string{
		StudentStatusAktif,
		StudentStatusCuti,
		StudentStatusNonAktif,
		StudentStatusMenungguUkom,
		StudentStatusMbkm,
	}
}

func MomentaryStudentStatus() []string {
	return []string{
		StudentStatusKeluar,
		StudentStatusLulus,
		StudentStatusDropOut,
	}
}

func ValidStudentStatus() []string {
	return []string{
		StudentStatusAktif,
		StudentStatusCuti,
		StudentStatusDropOut,
		StudentStatusKeluar,
		StudentStatusLulus,
		StudentStatusMbkm,
		StudentStatusMenungguUkom,
		StudentStatusNonAktif,
	}
}

func ManualEditableStatus() []string {
	return []string{
		StudentStatusAktif,
		StudentStatusKeluar,
		StudentStatusLulus,
		StudentStatusNonAktif,
		StudentStatusDropOut,
		StudentStatusMenungguUkom,
		StudentStatusMbkm,
	}
}
