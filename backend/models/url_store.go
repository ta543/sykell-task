package models

import (
	"database/sql"
	"time"
)

func CreateURL(u *URL) error {
	now := time.Now()
	result, err := DB.Exec(`
        INSERT INTO urls (address, title, html_version, h1_count, h2_count, h3_count,
        internal_links, external_links, broken_links, has_login_form, status, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `, u.Address, u.Title, u.HTMLVersion, u.H1Count, u.H2Count, u.H3Count,
		u.InternalLinks, u.ExternalLinks, u.BrokenLinks, u.HasLoginForm,
		u.Status, now, now)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

func UpdateURL(u *URL) error {
	u.UpdatedAt = time.Now()
	_, err := DB.Exec(`
        UPDATE urls SET title=?, html_version=?, h1_count=?, h2_count=?, h3_count=?, internal_links=?, external_links=?, broken_links=?, has_login_form=?, status=?, updated_at=?
        WHERE id=?
    `, u.Title, u.HTMLVersion, u.H1Count, u.H2Count, u.H3Count, u.InternalLinks, u.ExternalLinks, u.BrokenLinks, u.HasLoginForm, u.Status, u.UpdatedAt, u.ID)

	return err
}

func AddBrokenLinks(urlID int64, links []BrokenLink) error {
        return CreateBrokenLinks(urlID, links)
}

func GetBrokenLinksForURL(urlID int64) ([]BrokenLink, error) {
        return GetBrokenLinks(urlID)
}

func GetURLs(limit, offset int, search, status, html string) ([]URL, error) {
        rows, err := DB.Query(`
        SELECT id, address, title, html_version, h1_count, h2_count, h3_count,
        internal_links, external_links, broken_links, has_login_form, status, created_at, updated_at
        FROM urls
        WHERE (address LIKE ? OR title LIKE ?) AND (?='' OR status=?) AND (?='' OR html_version=?)
        ORDER BY id DESC LIMIT ? OFFSET ?
    `, "%"+search+"%", "%"+search+"%", status, status, html, html, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []URL
	for rows.Next() {
		var u URL
		err := rows.Scan(&u.ID, &u.Address, &u.Title, &u.HTMLVersion, &u.H1Count, &u.H2Count, &u.H3Count,
			&u.InternalLinks, &u.ExternalLinks, &u.BrokenLinks, &u.HasLoginForm, &u.Status, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, u)
	}
	return list, nil
}

func GetURL(id int64) (*URL, error) {
	var u URL
	err := DB.QueryRow(`
        SELECT id, address, title, html_version, h1_count, h2_count, h3_count,
        internal_links, external_links, broken_links, has_login_form, status, created_at, updated_at
        FROM urls WHERE id=?
    `, id).Scan(&u.ID, &u.Address, &u.Title, &u.HTMLVersion, &u.H1Count, &u.H2Count, &u.H3Count,
		&u.InternalLinks, &u.ExternalLinks, &u.BrokenLinks, &u.HasLoginForm, &u.Status, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func DeleteURL(id int64) error {
	_, err := DB.Exec(`DELETE FROM urls WHERE id=?`, id)
	return err
}
