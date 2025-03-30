-- name: GetPerson :one
SELECT *
FROM persons
WHERE id = ?;

-- name: GetPersonByEmail :one
SELECT *
FROM persons
WHERE email = ?;

-- name: ListPersons :many
SELECT *
FROM persons
ORDER BY full_name;

-- name: CreatePerson :one
INSERT INTO persons (
  email,
  full_name,
  course_id
) VALUES (
  ?,
  ?,
  ?
)
RETURNING *;

-- name: UpdatePerson :execresult
UPDATE persons
SET email = ?,
    full_name = ?,
    course_id = ?
WHERE id = ?;

-- name: DeletePerson :execresult
DELETE FROM persons
WHERE id = ?;

-- name: UpdatePersonCourse :execresult
UPDATE persons
SET course_id = ?
WHERE id = ?;