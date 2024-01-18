package reminders

import (
	db2 "ahti/app/common/db"
	"ahti/app/config"
	"database/sql"
	"github.com/jasonlvhit/gocron"
	"go.uber.org/zap"
)

type service struct {
	logger *zap.SugaredLogger
	db     *sql.DB
}

func newService(logger *zap.SugaredLogger) *service {

	svc := service{
		logger: logger,
		db:     db2.NewDbInstance(config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, config.DatabaseName),
	}

	err := svc.runMigration()
	if err != nil {
		panic(err)
	}
	svc.logger.Info("migration done")

	svc.start()

	return &svc

}

func (s *service) start() {
	s.checkDBforRemindersAndRemindAboutTheTasks()
	s.logger.Info("stating reminder service")
	go func() {
		_ = gocron.Every(1).Minutes().Do(s.checkDBforRemindersAndRemindAboutTheTasks)
		gocron.Start()
	}()
}

func (s *service) checkDBforRemindersAndRemindAboutTheTasks() {
	//s.logger.Info("checking for reminders")
	list, err := s.getJobsToBeTriggeredNow()
	if err != nil {
		s.logger.Error(err)
		return
	}
	for _, reminder := range list {
		s.logger.Info(reminder.Description)
	}
	//s.logger.Infof("reminded total : %v tasks ", len(list))
}
