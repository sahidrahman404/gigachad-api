package ent

import (
	"context"
	"fmt"

	"github.com/sahidrahman404/gigachad-api/ent/equipment"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/exercisetype"
	"github.com/sahidrahman404/gigachad-api/ent/musclesgroup"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/token"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// prefixMap maps PULID prefixes to table names.
var prefixMap = map[pksuid.ID]string{
	"UR": user.Table,
	"TK": token.Table,
	"EX": exercise.Table,
	"MG": musclesgroup.Table,
	"ET": exercisetype.Table,
	"EQ": equipment.Table,
	"RO": routine.Table,
	"RE": routineexercise.Table,
	"WO": workout.Table,
	"WL": workoutlog.Table,
}

// IDToType maps a pulid.ID to the underlying table.
func IDToType(ctx context.Context, id pksuid.ID) (string, error) {
	if len(id) < 2 {
		return "", fmt.Errorf("IDToType: id too short")
	}
	prefix := id[:2]
	typ := prefixMap[prefix]
	if typ == "" {
		return "", fmt.Errorf("IDToType: could not map prefix '%s' to a type", prefix)
	}
	return typ, nil
}
