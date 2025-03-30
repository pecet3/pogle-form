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

-- name: CreatePerson :execresult
INSERT INTO persons (
  email,
  full_name,
  chosen_course_id
) VALUES (
  ?,
  ?,
  ?
);

-- name: UpdatePerson :execresult
UPDATE persons
SET email = ?,
    full_name = ?,
    chosen_course_id = ?
WHERE id = ?;

-- name: DeletePerson :execresult
DELETE FROM persons
WHERE id = ?;

-- name: UpdatePersonChosenCourse :execresult
UPDATE persons
SET chosen_course_id = ?
WHERE id = ?;