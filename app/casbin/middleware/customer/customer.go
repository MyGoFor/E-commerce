package customer

import "github.com/casbin/casbin/v2"

func Check(e *casbin.CachedEnforcer, sub, obj, act string) bool {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		return true
	} else {
		return false
	}
}

func AddCustomer(e *casbin.CachedEnforcer, sub string) error {
	_, err := e.AddRoleForUser(sub, "customer")
	if err != nil {
		return err
	}
	return nil
}

func DeleteCustomer(e *casbin.CachedEnforcer, sub string) error {
	_, err := e.RemoveGroupingPolicy(sub, "customer")
	if err != nil {
		return err
	}
	return nil
}
