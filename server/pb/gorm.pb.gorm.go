package pb

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strings "strings"
	time "time"
)

type ExampleORM struct {
	CratedAt  *time.Time
	DeletedAt *time.Time
	Id        uint64 `gorm:"primary_key;not null"`
	UpdatedAt *time.Time
}

// TableName overrides the default tablename generated by GORM
func (ExampleORM) TableName() string {
	return "example_table"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Example) ToORM(ctx context.Context) (ExampleORM, error) {
	to := ExampleORM{}
	var err error
	if prehook, ok := interface{}(m).(ExampleWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.CratedAt != nil {
		t := m.CratedAt.AsTime()
		to.CratedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if m.DeletedAt != nil {
		t := m.DeletedAt.AsTime()
		to.DeletedAt = &t
	}
	if posthook, ok := interface{}(m).(ExampleWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ExampleORM) ToPB(ctx context.Context) (Example, error) {
	to := Example{}
	var err error
	if prehook, ok := interface{}(m).(ExampleWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.CratedAt != nil {
		to.CratedAt = timestamppb.New(*m.CratedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if m.DeletedAt != nil {
		to.DeletedAt = timestamppb.New(*m.DeletedAt)
	}
	if posthook, ok := interface{}(m).(ExampleWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Example the arg will be the target, the caller the one being converted from

// ExampleBeforeToORM called before default ToORM code
type ExampleWithBeforeToORM interface {
	BeforeToORM(context.Context, *ExampleORM) error
}

// ExampleAfterToORM called after default ToORM code
type ExampleWithAfterToORM interface {
	AfterToORM(context.Context, *ExampleORM) error
}

// ExampleBeforeToPB called before default ToPB code
type ExampleWithBeforeToPB interface {
	BeforeToPB(context.Context, *Example) error
}

// ExampleAfterToPB called after default ToPB code
type ExampleWithAfterToPB interface {
	AfterToPB(context.Context, *Example) error
}

// DefaultCreateExample executes a basic gorm create call
func DefaultCreateExample(ctx context.Context, in *Example, db *gorm.DB) (*Example, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type ExampleORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExampleORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadExample(ctx context.Context, in *Example, db *gorm.DB) (*Example, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &ExampleORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := ExampleORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ExampleORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type ExampleORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExampleORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExampleORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteExample(ctx context.Context, in *Example, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&ExampleORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type ExampleORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExampleORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteExampleSet(ctx context.Context, in []*Example, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&ExampleORM{})).(ExampleORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&ExampleORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&ExampleORM{})).(ExampleORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type ExampleORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Example, *gorm.DB) (*gorm.DB, error)
}
type ExampleORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Example, *gorm.DB) error
}

// DefaultStrictUpdateExample clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateExample(ctx context.Context, in *Example, db *gorm.DB) (*Example, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateExample")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &ExampleORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ExampleORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type ExampleORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExampleORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExampleORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchExample executes a basic gorm update call with patch behavior
func DefaultPatchExample(ctx context.Context, in *Example, updateMask *field_mask.FieldMask, db *gorm.DB) (*Example, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Example
	var err error
	if hook, ok := interface{}(&pbObj).(ExampleWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadExample(ctx, &Example{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ExampleWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskExample(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ExampleWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateExample(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ExampleWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type ExampleWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Example, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ExampleWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Example, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ExampleWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Example, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ExampleWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Example, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetExample executes a bulk gorm update call with patch behavior
func DefaultPatchSetExample(ctx context.Context, objects []*Example, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Example, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Example, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchExample(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskExample patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskExample(ctx context.Context, patchee *Example, patcher *Example, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Example, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedCratedAt bool
	var updatedUpdatedAt bool
	var updatedDeletedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if !updatedCratedAt && strings.HasPrefix(f, prefix+"CratedAt.") {
			if patcher.CratedAt == nil {
				patchee.CratedAt = nil
				continue
			}
			if patchee.CratedAt == nil {
				patchee.CratedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CratedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CratedAt, patchee.CratedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CratedAt" {
			updatedCratedAt = true
			patchee.CratedAt = patcher.CratedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if !updatedDeletedAt && strings.HasPrefix(f, prefix+"DeletedAt.") {
			if patcher.DeletedAt == nil {
				patchee.DeletedAt = nil
				continue
			}
			if patchee.DeletedAt == nil {
				patchee.DeletedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"DeletedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.DeletedAt, patchee.DeletedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"DeletedAt" {
			updatedDeletedAt = true
			patchee.DeletedAt = patcher.DeletedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListExample executes a gorm list call
func DefaultListExample(ctx context.Context, db *gorm.DB) ([]*Example, error) {
	in := Example{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &ExampleORM{}, &Example{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []ExampleORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExampleORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Example{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ExampleORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExampleORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExampleORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]ExampleORM) error
}
