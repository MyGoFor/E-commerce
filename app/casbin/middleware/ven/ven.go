package ven

import "github.com/casbin/casbin/v2"

func Check(e *casbin.CachedEnforcer, sub string) bool {
	ok, _ := e.Enforce(sub, "vendor")
	if ok {
		return true
	} else {
		return false
	}
}

func AddVendor(e *casbin.CachedEnforcer, sub string) error {
	_, err := e.AddRoleForUser(sub, "vendor")
	if err != nil {
		return err
	}
	return nil
}

func DeleteVendor(e *casbin.CachedEnforcer, sub string) error {
	_, err := e.RemoveGroupingPolicy(sub, "vendor")
	if err != nil {
		return err
	}
	return nil
}
