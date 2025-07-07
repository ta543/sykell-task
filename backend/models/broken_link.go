package models

import "time"

// BrokenLink represents a single broken link for a URL.
type BrokenLink struct {
    ID         int64 `json:"id" db:"id"`
    URLID      int64 `json:"url_id" db:"url_id"`
    Link       string `json:"link" db:"link"`
    StatusCode int   `json:"status_code" db:"status_code"`
    CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

func CreateBrokenLinks(urlID int64, links []BrokenLink) error {
    for _, l := range links {
        _, err := DB.Exec(`INSERT INTO broken_links (url_id, link, status_code, created_at) VALUES (?, ?, ?, ?)`,
            urlID, l.Link, l.StatusCode, time.Now())
        if err != nil {
            return err
        }
    }
    return nil
}

func GetBrokenLinks(urlID int64) ([]BrokenLink, error) {
    rows, err := DB.Query(`SELECT id, url_id, link, status_code, created_at FROM broken_links WHERE url_id=?`, urlID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var list []BrokenLink
    for rows.Next() {
        var b BrokenLink
        if err := rows.Scan(&b.ID, &b.URLID, &b.Link, &b.StatusCode, &b.CreatedAt); err != nil {
            return nil, err
        }
        list = append(list, b)
    }
    return list, nil
}

