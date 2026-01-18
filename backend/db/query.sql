-- name: CreateUser :execresult
INSERT INTO User (first_name, second_name, email, isAdmin, isBanned)
VALUES (?, ?, ?, ?, ?);

-- name: CreateProject :execresult
INSERT INTO Project (name, code)
VALUES (?, ?);

-- name: CreateUserProject :execresult
INSERT INTO User_Project (user_id, project_id)
VALUES (?, ?);

-- name: CreateStation :execresult
INSERT INTO Station (name, project_id)
VALUES (?, ?);

-- name: CreateState :execresult
INSERT INTO State (final_state, start_date, end_date, station_id)
VALUES (?, ?, ?, ?);

-- name: ListProjects :many
SELECT * FROM Project;

-- name: GetProject :one
SELECT * FROM Project WHERE id = ?;

-- name: ListUsers :many
SELECT * FROM User;

-- name: GetUser :one
SELECT * FROM User WHERE id = ?;

-- name: GetState :one
SELECT * FROM State WHERE id = ?;

-- Get all stations associated with a specific project
-- name: ListStationsByProject :many
SELECT s.id, s.name, p.name as project_name
FROM Station s
JOIN Project p ON s.project_id = p.id
WHERE p.id = ?;

-- Get the history of states for a specific station, ordered by start date
-- name: ListStationHistory :many
SELECT st.final_state, st.start_date, st.end_date
FROM State st
WHERE st.station_id = ?
ORDER BY st.start_date DESC;

-- Find the current active state for a station (assuming null end_date means active)
-- name: GetActiveStationState :one
SELECT * FROM State
WHERE station_id = ? AND end_date IS NULL;

-- List all users assigned to a specific project
-- name: ListUsersByProject :many
SELECT u.first_name, u.second_name, u.email
FROM User u
JOIN User_Project up ON u.id = up.user_id
WHERE up.project_id = ?;

-- Count the number of stations per project
-- name: CountStationsPerProject :many
SELECT p.name, COUNT(s.id) as station_count
FROM Project p
LEFT JOIN Station s ON p.id = s.project_id
GROUP BY p.id;

-- Find projects that have no users assigned
-- name: ListProjectsWithoutUsers :many
SELECT p.name
FROM Project p
LEFT JOIN User_Project up ON p.id = up.project_id
WHERE up.user_id IS NULL;
