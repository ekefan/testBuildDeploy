// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: tripBookings.sql

package db

import (
	"context"
)

const createTripBooking = `-- name: CreateTripBooking :one
INSERT INTO trip_bookings(
    trip_owner,
    booking_type,
    booking_details
) VALUES (
    $1, $2, $3
) RETURNING id, trip_owner, booking_type, booking_details
`

type CreateTripBookingParams struct {
	TripOwner      int64  `json:"trip_owner"`
	BookingType    string `json:"booking_type"`
	BookingDetails string `json:"booking_details"`
}

func (q *Queries) CreateTripBooking(ctx context.Context, arg CreateTripBookingParams) (TripBooking, error) {
	row := q.queryRow(ctx, q.createTripBookingStmt, createTripBooking, arg.TripOwner, arg.BookingType, arg.BookingDetails)
	var i TripBooking
	err := row.Scan(
		&i.ID,
		&i.TripOwner,
		&i.BookingType,
		&i.BookingDetails,
	)
	return i, err
}

const deleteTripBooking = `-- name: DeleteTripBooking :exec
DELETE FROM trip_bookings WHERE trip_owner = &1
`

func (q *Queries) DeleteTripBooking(ctx context.Context) error {
	_, err := q.exec(ctx, q.deleteTripBookingStmt, deleteTripBooking)
	return err
}

const getTripBookingUpdate = `-- name: GetTripBookingUpdate :one
SELECT id, trip_owner, booking_type, booking_details FROM trip_bookings
WHERE trip_owner = $1
LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetTripBookingUpdate(ctx context.Context, tripOwner int64) (TripBooking, error) {
	row := q.queryRow(ctx, q.getTripBookingUpdateStmt, getTripBookingUpdate, tripOwner)
	var i TripBooking
	err := row.Scan(
		&i.ID,
		&i.TripOwner,
		&i.BookingType,
		&i.BookingDetails,
	)
	return i, err
}

const listTripBooking = `-- name: ListTripBooking :many
SELECT FROM trip_bookings
WHERE trip_owner = $1
LIMIT $2
OFFSET $3
`

type ListTripBookingParams struct {
	TripOwner int64 `json:"trip_owner"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

type ListTripBookingRow struct {
}

func (q *Queries) ListTripBooking(ctx context.Context, arg ListTripBookingParams) ([]ListTripBookingRow, error) {
	rows, err := q.query(ctx, q.listTripBookingStmt, listTripBooking, arg.TripOwner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListTripBookingRow{}
	for rows.Next() {
		var i ListTripBookingRow
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
