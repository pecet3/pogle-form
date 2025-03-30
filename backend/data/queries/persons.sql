-- name: CreatePerson :execresult
INSERT INTO persons (
  email,
  full_name
) VALUES (
  ?, ?
);

-- name: GetPerson :one
SELECT * FROM persons
WHERE id = ? LIMIT 1;

-- name: GetPersonByEmail :one
SELECT * FROM persons
WHERE email = ? LIMIT 1;

-- name: ListPersons :many
SELECT * FROM persons
ORDER BY full_name;

-- name: UpdatePerson :execresult
UPDATE persons SET
  email = ?,
  full_name = ?
WHERE id = ?;

-- name: DeletePerson :execresult
DELETE FROM persons
WHERE id = ?;