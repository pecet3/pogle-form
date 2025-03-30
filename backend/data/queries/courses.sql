-- courses table queries

-- name: GetCourse :one
SELECT *
FROM courses
WHERE id = ?;

-- name: GetAllCourses :many
SELECT *
FROM courses
ORDER BY name;

-- name: CreateCourse :execresult
INSERT INTO courses (
  name,
  max_persons
) VALUES (
  ?,
  ?
);

-- name: UpdateCourse :execresult
UPDATE courses
SET name = ?,
    max_persons = ?
WHERE id = ?;

-- name: DeleteCourse :execresult
DELETE FROM courses
WHERE id = ?;


-- chosen_courses table queries

-- name: GetChosenCourse :one
SELECT *
FROM chosen_courses
WHERE id = ?;

-- name: ListChosenCoursesByPersonID :many
SELECT *
FROM chosen_courses
WHERE person_id = ?
ORDER BY id;

-- name: ListChosenCoursesByCourseID :many
SELECT *
FROM chosen_courses
WHERE course_id = ?
ORDER BY id;

-- name: CreateChosenCourse :execresult
INSERT INTO chosen_courses (
  person_id,
  course_id
) VALUES (
  ?,
  ?
);

-- name: DeleteChosenCourse :execresult
DELETE FROM chosen_courses
WHERE id = ?;

-- name: DeleteChosenCoursesByPersonID :execresult
DELETE FROM chosen_courses
WHERE person_id = ?;

-- name: DeleteChosenCoursesByCourseID :execresult
DELETE FROM chosen_courses
WHERE course_id = ?;


-- chosen_reserved_courses table queries

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
  COALESCE(COUNT(cc.person_id), 0) AS number_of_persons
FROM
  courses c
LEFT JOIN
  chosen_courses cc ON c.id = cc.course_id
WHERE
  c.id = ?;