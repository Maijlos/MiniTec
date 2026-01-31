-- name: CreateProject :execresult
INSERT INTO Project (name, code)
VALUES (?, ?);

-- name: UpdateProject :exec
UPDATE Project
SET name = ?, code = ?
WHERE id = ?;

-- name: CreateStation :execresult
INSERT INTO Station (name, project_id)
VALUES (?, ?);

-- name: GetStationId :one
SELECT id FROM Station WHERE name = ? AND project_id = ?;

-- name: CreateState :execresult
INSERT INTO State (final_state, start_date, end_date, station_id)
VALUES (?, ?, ?, ?);

-- name: ListProjects :many
SELECT * FROM Project 
ORDER BY code 
LIMIT ? OFFSET ?;

-- name: GetProject :one
SELECT * FROM Project WHERE id = ?;

-- name: GetProjectByCode :one
SELECT * FROM Project WHERE code = ?;

-- name: DeleteProject :execresult
DELETE FROM Project WHERE id = ?;

-- name: GetState :one
SELECT * FROM State WHERE id = ?;

-- Get all stations associated with a specific project
-- name: ListStationsByProject :many
SELECT * FROM Station WHERE project_id = ?;

-- Get all states associated with a specific station
-- name: ListStationsByStation :many
SELECT * FROM State WHERE station_id = ?
ORDER BY State.start_date ASC;
