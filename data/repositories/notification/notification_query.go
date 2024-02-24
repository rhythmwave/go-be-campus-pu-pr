package notification

const (
	bulkCreateQuery = `
		INSERT INTO notifications (
			user_id,
			title,
			body,
			link_url,
			photo_path,
			photo_path_type,
			expired_at,
			id_type,
			type,
			data
		) VALUES (
			:user_id,
			:title,
			:body,
			:link_url,
			:photo_path,
			:photo_path_type,
			:expired_at,
			:id_type,
			:type,
			:data
		)
	`
)
