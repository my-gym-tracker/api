package api

type Error int

const (
	ErrorBadRequest          Error = 400
	ErrorUnauthorized        Error = 401
	ErrorNotFound            Error = 404
	ErrorConflict            Error = 409
	ErrorInternalServerError Error = 500
)

type ErrorRes struct {
	Code   Error  `json:"code"`
	Reason string `json:"reason"`
}

type UserCredentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type RegisterReq struct {
	UserCredentials
}

type RegisterRes struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}

type LoginReq struct {
	UserCredentials
}

type LoginRes struct {
	Token  string `json:"token"`
	UserId int    `json:"userId"`
}

type User struct {
	UserId int `json:"userId"`
	UserCredentials
}

type GetUserReq struct {
	UserId int `json:"userId"`
}

type CreateUserReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UpdateUserReq struct {
	UserId       int    `json:"userId"`
	UserName     string `json:"userName"`
	UserPassword string `json:"userPassword"`
}

type DeleteUserReq struct {
	UserId int `json:"userId"`
}

type GetUserRes struct {
	User
}

type CreateUserRes struct {
	User
}

type UpdateUserRes struct {
	User
}

type DeleteUserRes struct{}

type Workout struct {
	WorkoutId int    `json:"workoutId"`
	UserId    int    `json:"userId"`
	Date      string `json:"date"`
}

type GetWorkoutReq struct {
	WorkoutId int `json:"workoutId"`
}

type CreateWorkoutReq struct {
	UserId int    `json:"userId"`
	Date   string `json:"date"`
}

type UpdateWorkoutReq struct {
	WorkoutId int    `json:"workoutId"`
	UserId    int    `json:"userId"`
	Date      string `json:"date"`
}

type DeleteWorkoutReq struct {
	WorkoutId int `json:"workoutId"`
}

type GetWorkoutRes struct {
	Workout
}

type CreateWorkoutRes struct {
	Workout
}

type UpdateWorkoutRes struct {
	Workout
}

type DeleteWorkoutRes struct{}

type Exercise struct {
	ExerciseId  int    `json:"exerciseId"`
	UserId      int    `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetExerciseReq struct {
	ExerciseId int `json:"exerciseId"`
}

type CreateExerciseReq struct {
	UserId              int    `json:"userId"`
	ExerciseName        string `json:"exerciseName"`
	ExerciseDescription string `json:"exerciseDescription"`
}

type UpdateExerciseReq struct {
	ExerciseId          int    `json:"exerciseId"`
	ExerciseName        string `json:"exerciseName"`
	ExerciseDescription string `json:"exerciseDescription"`
}

type DeleteExerciseReq struct {
	ExerciseId int `json:"exerciseId"`
}

type GetExerciseRes struct {
	Exercise
}

type CreateExerciseRes struct {
	Exercise
}

type UpdateExerciseRes struct {
	Exercise
}

type DeleteExerciseRes struct{}

type Set struct {
	SetId        int     `json:"setId"`
	WorkoutId    int     `json:"workoutId"`
	ExerciseId   int     `json:"exerciseId"`
	Index        int     `json:"index"`
	Date         string  `json:"date"`
	Reps         int     `json:"reps"`
	WeightType   int     `json:"weightType"`
	WeightAmount float32 `json:"weightAmount"`
}

type GetSetReq struct {
	SetId int `json:"setId"`
}

type CreateSetReq struct {
	WorkoutId    int     `json:"workoutId"`
	ExerciseId   int     `json:"exerciseId"`
	Index        int     `json:"index"`
	Date         string  `json:"date"`
	Reps         int     `json:"reps"`
	WeightType   int     `json:"weightType"`
	WeightAmount float32 `json:"weightAmount"`
}

type UpdateSetReq struct {
	SetId        int     `json:"setId"`
	ExerciseId   int     `json:"exerciseId"`
	Index        int     `json:"index"`
	Date         string  `json:"date"`
	Reps         int     `json:"reps"`
	WeightType   int     `json:"weightType"`
	WeightAmount float32 `json:"weightAmount"`
}

type DeleteSetReq struct {
	SetId int `json:"setId"`
}

type GetSetRes struct {
	Set
}

type CreateSetRes struct {
	Set
}

type UpdateSetRes struct {
	Set
}

type DeleteSetRes struct{}
