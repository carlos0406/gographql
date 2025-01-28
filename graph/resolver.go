package graph

import internaldatabase "graphql/internal/database"

type Resolver struct {
	CategoryDB *internaldatabase.Category
	CourseDB   *internaldatabase.Course
}
