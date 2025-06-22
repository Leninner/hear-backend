-- name: CreateClassroom :one
INSERT INTO classrooms (
    name,
    building,
    floor,
    capacity,
    faculty_id,
    location_lat,
    location_lng
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetClassroomByID :one
SELECT * FROM classrooms
WHERE id = $1 LIMIT 1;

-- name: GetClassroomByName :one
SELECT * FROM classrooms
WHERE name = $1 LIMIT 1;

-- name: GetClassroomsByFacultyID :many
SELECT * FROM classrooms
WHERE faculty_id = $1
ORDER BY name;

-- name: GetClassroomsByBuilding :many
SELECT * FROM classrooms
WHERE building = $1;

-- name: GetClassroomsByCapacity :many
SELECT * FROM classrooms
WHERE capacity >= $1;

-- name: GetAll :many
SELECT * FROM classrooms
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetNearbyClassrooms :many
SELECT *,
    earth_distance(ll_to_earth(location_lat, location_lng), ll_to_earth($1, $2)) as distance
FROM classrooms
WHERE earth_box(ll_to_earth($1, $2), $3) @> ll_to_earth(location_lat, location_lng)
    AND earth_distance(ll_to_earth(location_lat, location_lng), ll_to_earth($1, $2)) <= $3
ORDER BY distance;

-- name: UpdateClassroom :one
UPDATE classrooms
SET
    name = COALESCE($2, name),
    building = COALESCE($3, building),
    floor = COALESCE($4, floor),
    capacity = COALESCE($5, capacity),
    faculty_id = COALESCE($6, faculty_id),
    location_lat = COALESCE($7, location_lat),
    location_lng = COALESCE($8, location_lng),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteClassroom :exec
DELETE FROM classrooms
WHERE id = $1; 