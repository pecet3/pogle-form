-- name: GetCourse :one
SELECT *
FROM courses
WHERE id = ?;

-- name: ListCourses :many
SELECT *
FROM courses
ORDER BY name;

-- name: CreateCourse :one
INSERT INTO courses (
  name,
  max_persons
) VALUES (
  ?,
  ?
)
RETURNING *;

-- name: UpdateCourse :execresult
UPDATE courses
SET name = ?,
    max_persons = ?
WHERE id = ?;

-- name: DeleteCourse :execresult
DELETE FROM courses
WHERE id = ?;



-- name: GetChosenReservedCourse :one
SELECT *
FROM chosen_reserved_courses
WHERE id = ?;

-- name: ListChosenReservedCoursesByPersonID :many
SELECT *
FROM chosen_reserved_courses
WHERE person_id = ?
ORDER BY id;

-- name: ListChosenReservedCoursesByCourseID :many
SELECT *
FROM chosen_reserved_courses
WHERE course_id = ?
ORDER BY id;

-- name: CreateChosenReservedCourse :execresult
INSERT INTO chosen_reserved_courses (
  person_id,
  course_id
) VALUES (
  ?,
  ?
);

-- name: DeleteChosenReservedCourse :execresult
DELETE FROM chosen_reserved_courses
WHERE id = ?;

-- name: DeleteChosenReservedCoursesByPersonID :execresult
DELETE FROM chosen_reserved_courses
WHERE person_id = ?;

-- name: DeleteChosenReservedCoursesByCourseID :execresult
DELETE FROM chosen_reserved_courses
WHERE course_id = ?;

-- name: GetNumberOfPersonsInCourse :one
SELECT
  COUNT(p.id) AS number_of_persons
FROM
  courses c
LEFT JOIN
  persons p ON c.id = p.course_id
WHERE
  c.id = ?;