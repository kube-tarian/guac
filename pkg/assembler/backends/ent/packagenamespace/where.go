// Code generated by ent, DO NOT EDIT.

package packagenamespace

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldLTE(FieldID, id))
}

// PackageID applies equality check predicate on the "package_id" field. It's identical to PackageIDEQ.
func PackageID(v int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldEQ(FieldPackageID, v))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldEQ(FieldNamespace, v))
}

// PackageIDEQ applies the EQ predicate on the "package_id" field.
func PackageIDEQ(v int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldEQ(FieldPackageID, v))
}

// PackageIDNEQ applies the NEQ predicate on the "package_id" field.
func PackageIDNEQ(v int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldNEQ(FieldPackageID, v))
}

// PackageIDIn applies the In predicate on the "package_id" field.
func PackageIDIn(vs ...int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldIn(FieldPackageID, vs...))
}

// PackageIDNotIn applies the NotIn predicate on the "package_id" field.
func PackageIDNotIn(vs ...int) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldNotIn(FieldPackageID, vs...))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.FieldContainsFold(FieldNamespace, v))
}

// HasPackage applies the HasEdge predicate on the "package" edge.
func HasPackage() predicate.PackageNamespace {
	return predicate.PackageNamespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PackageTable, PackageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPackageWith applies the HasEdge predicate on the "package" edge with a given conditions (other predicates).
func HasPackageWith(preds ...predicate.PackageType) predicate.PackageNamespace {
	return predicate.PackageNamespace(func(s *sql.Selector) {
		step := newPackageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNames applies the HasEdge predicate on the "names" edge.
func HasNames() predicate.PackageNamespace {
	return predicate.PackageNamespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NamesTable, NamesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamesWith applies the HasEdge predicate on the "names" edge with a given conditions (other predicates).
func HasNamesWith(preds ...predicate.PackageName) predicate.PackageNamespace {
	return predicate.PackageNamespace(func(s *sql.Selector) {
		step := newNamesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PackageNamespace) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PackageNamespace) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PackageNamespace) predicate.PackageNamespace {
	return predicate.PackageNamespace(sql.NotPredicates(p))
}
