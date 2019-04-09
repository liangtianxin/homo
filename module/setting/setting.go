//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package setting

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"sync"
)

var (
	BeepSpeakerInited bool
	DebugMode         bool
	IntentOnlyMode    bool
	VoicePlayMutex    sync.Mutex
)

func init() {
	BeepSpeakerInited = false
	DebugMode = false
	IntentOnlyMode = false
}

func NewContext() {
	//Initial portaudio
	/*err := portaudio.Initialize()
	if err != nil {
		logrus.Fatalf("Initial portaudio failed: %s", err.Error())
	}*/
}

func Terminal(c *cli.Context) error {
	logrus.Infof("退出，开始结束进程...")
	/*err := portaudio.Terminate()
	if err != nil {
		logrus.Warnf("Close portaudio failed", err.Error())
		return err
	}*/
	return nil
}