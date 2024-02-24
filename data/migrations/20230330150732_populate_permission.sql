-- +goose Up
-- +goose StatementBegin

DO $$
DECLARE adminId uuid;
BEGIN
  SELECT id INTO adminId FROM admins WHERE role_id IS NULL LIMIT 1;

  IF adminId IS NULL THEN
    RAISE EXCEPTION 'root admin is not found';
  END IF;

  INSERT INTO permissions (name, created_by) VALUES
    ('Aktivitas Mahasiswa', adminId),
    ('Data Referensi', adminId),
    ('Dosen', adminId),
    ('Evaluasi', adminId),
    ('Hasil Studi', adminId),
    ('Input Nilai', adminId),
    ('Jadwal', adminId),
    ('Jatah Sks', adminId),
    ('Jenis Nilai', adminId),
    ('Kelulusan', adminId),
    ('Kemahasiswaan', adminId),
    ('Komponen Nilai', adminId),
    ('Kuisioner', adminId),
    ('Kurikulum', adminId),
    ('Laporan', adminId),
    ('Manajemen Ruang', adminId),
    ('Menunggu UKOM', adminId),
    ('Nilai', adminId),
    ('Pejabat', adminId),
    ('Penyelenggara MBKM', adminId),
    ('Persetujuan KRS', adminId),
    ('Peserta Kelas', adminId),
    ('Peserta MBKM', adminId),
    ('Portal Akademik', adminId),
    ('Predikat Kelulusan', adminId),
    ('Program Studi', adminId),
    ('Rekap Nilai', adminId),
    ('Rencana Studi', adminId),
    ('SKPI', adminId),
    ('Semester', adminId),
    ('Status Registrasi', adminId),
    ('Transkrip', adminId),
    ('Tugas Akhir', adminId),
    ('Virtual Class', adminId);
END;
$$;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DELETE FROM permissions;

-- +goose StatementEnd
