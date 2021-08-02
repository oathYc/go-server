package boot

import (
	"github.com/pkg/errors"
)

func Bootstrap(configPath string) (err error) {

	if err = bootConfig(configPath); nil != err {
		return errors.Wrap(err, "bootConfig")
	}

	// 初始化日志
	if err = bootLogger(); err != nil {
		return errors.Wrap(err, "bootLogger")
	}

	if err = bootMysql(); nil != err {
		return errors.Wrap(err, "bootMysql")
	}

	if err = InitRedisPool(); nil != err {
		return errors.Wrap(err, "bootRedis")
	}
	return nil
}
