package model

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// Corporation represents a row from 'Attractech.corporation'.
type Corporation struct {
	ID        uint64    `json:"id"`         // id
	UpdatedAt time.Time `json:"updated_at"` // updated_at
	CreatedAt time.Time `json:"created_at"` // created_at
	PublicID  string    `json:"public_id"`  // public_id
	Name      string    `json:"name"`       // name
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Corporation exists in the database.
func (c *Corporation) Exists() bool {
	return c._exists
}

// Deleted returns true when the Corporation has been marked for deletion from
// the database.
func (c *Corporation) Deleted() bool {
	return c._deleted
}

// Insert inserts the Corporation to the database.
func (c *Corporation) Insert(ctx context.Context, db DB) error {
	switch {
	case c._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case c._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO Attractech.corporation (` +
		`updated_at, created_at, public_id, name` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, c.UpdatedAt, c.CreatedAt, c.PublicID, c.Name)
	res, err := db.ExecContext(ctx, sqlstr, c.UpdatedAt, c.CreatedAt, c.PublicID, c.Name)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	c.ID = uint64(id)
	// set exists
	c._exists = true
	return nil
}

// Update updates a Corporation in the database.
func (c *Corporation) Update(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case c._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE Attractech.corporation SET ` +
		`updated_at = ?, created_at = ?, public_id = ?, name = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, c.UpdatedAt, c.CreatedAt, c.PublicID, c.Name, c.ID)
	if _, err := db.ExecContext(ctx, sqlstr, c.UpdatedAt, c.CreatedAt, c.PublicID, c.Name, c.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Corporation to the database.
func (c *Corporation) Save(ctx context.Context, db DB) error {
	if c.Exists() {
		return c.Update(ctx, db)
	}
	return c.Insert(ctx, db)
}

// Upsert performs an upsert for Corporation.
func (c *Corporation) Upsert(ctx context.Context, db DB) error {
	switch {
	case c._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO Attractech.corporation (` +
		`id, updated_at, created_at, public_id, name` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`updated_at = VALUES(updated_at), created_at = VALUES(created_at), public_id = VALUES(public_id), name = VALUES(name)`
	// run
	logf(sqlstr, c.ID, c.UpdatedAt, c.CreatedAt, c.PublicID, c.Name)
	if _, err := db.ExecContext(ctx, sqlstr, c.ID, c.UpdatedAt, c.CreatedAt, c.PublicID, c.Name); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Delete deletes the Corporation from the database.
func (c *Corporation) Delete(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return nil
	case c._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM Attractech.corporation ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, c.ID)
	if _, err := db.ExecContext(ctx, sqlstr, c.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// CorporationByID retrieves a row from 'Attractech.corporation' as a Corporation.
//
// Generated from index 'corporation_id_pkey'.
func CorporationByID(ctx context.Context, db DB, id uint64) (*Corporation, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, updated_at, created_at, public_id, name ` +
		`FROM Attractech.corporation ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	c := Corporation{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&c.ID, &c.UpdatedAt, &c.CreatedAt, &c.PublicID, &c.Name); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}

// CorporationByName retrieves a row from 'Attractech.corporation' as a Corporation.
//
// Generated from index 'name'.
func CorporationByName(ctx context.Context, db DB, name string) (*Corporation, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, updated_at, created_at, public_id, name ` +
		`FROM Attractech.corporation ` +
		`WHERE name = ?`
	// run
	logf(sqlstr, name)
	c := Corporation{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, name).Scan(&c.ID, &c.UpdatedAt, &c.CreatedAt, &c.PublicID, &c.Name); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}

// CorporationByPublicID retrieves a row from 'Attractech.corporation' as a Corporation.
//
// Generated from index 'public_id'.
func CorporationByPublicID(ctx context.Context, db DB, publicID string) (*Corporation, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, updated_at, created_at, public_id, name ` +
		`FROM Attractech.corporation ` +
		`WHERE public_id = ?`
	// run
	logf(sqlstr, publicID)
	c := Corporation{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, publicID).Scan(&c.ID, &c.UpdatedAt, &c.CreatedAt, &c.PublicID, &c.Name); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}
