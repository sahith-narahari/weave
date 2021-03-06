package termdeposit

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/gconf"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/orm"
)

func init() {
	migration.MustRegister(1, &Configuration{}, migration.NoModification)
}

var _ orm.Model = (*Configuration)(nil)

func (c *Configuration) Validate() error {
	var errs error
	errs = errors.AppendField(errs, "Metadata", c.Metadata.Validate())
	errs = errors.AppendField(errs, "Owner", c.Owner.Validate())
	errs = errors.AppendField(errs, "Admin", c.Admin.Validate())
	if len(c.Bonuses) == 0 {
		errs = errors.AppendField(errs, "Bonuses", errors.ErrEmpty)
	}
	const maxBaseRates = 100 // Arbitrary limit to avoid huge data set.
	if len(c.BaseRates) > maxBaseRates {
		errs = errors.AppendField(errs, "BaseRates",
			errors.Wrapf(errors.ErrInput, "at most %d elements can be provided", maxBaseRates))
	}
	if hasDuplicates(c.BaseRates) {
		errs = errors.AppendField(errs, "BaseRates", errors.ErrDuplicate)
	}
	return errs
}

func hasDuplicates(rates []CustomRate) bool {
	addrs := make(map[string]struct{})
	for _, r := range rates {
		a := r.Address.String()
		if _, ok := addrs[a]; ok {
			return true
		}
		addrs[a] = struct{}{}
	}
	return false
}

// bestDepositBonus returns the best available for given period deposit bonus
// value. This function returns nil if no match was found.
func bestDepositBonus(bonuses []DepositBonus, duration weave.UnixDuration) *DepositBonus {
	var best *DepositBonus
	for _, b := range bonuses {
		if b.LockinPeriod > duration {
			continue
		}
		if best == nil || b.Bonus.Compare(best.Bonus) > 0 {
			best = &DepositBonus{LockinPeriod: b.LockinPeriod, Bonus: b.Bonus}
		}
	}
	return best
}

func loadConf(db gconf.Store) (Configuration, error) {
	var conf Configuration
	if err := gconf.Load(db, "termdeposit", &conf); err != nil {
		return conf, errors.Wrap(err, "load configuration")
	}
	return conf, nil
}
