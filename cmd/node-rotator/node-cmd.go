package main

import (
	rotatorModel "github.com/mattermost/rotator/model"
	"github.com/mattermost/rotator/rotator"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type rotateOptions struct {
	cluster              string
	maxScaling           uint
	rotateMasters        bool
	rotateWorkers        bool
	maxDrainRetries      uint
	evictGracePeriod     uint
	waitBetweenRotations uint
	waitBetweenDrains    uint
}

var (
	whatamidoing string
)

func newRotateCmd() *cobra.Command {
	o := rotateOptions{}
	cmd := &cobra.Command{
		Use:   "roll",
		Short: "Rollover nodes",
		Long:  `Applying latest AMI that exists on launch template or launch configuration and rolls out new nodes`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logger.WithField("cluster", o.cluster)

			clusterRotator := rotatorModel.Cluster{
				ClusterID:            o.cluster,
				MaxScaling:           int(o.maxScaling),
				RotateMasters:        o.rotateMasters,
				RotateWorkers:        o.rotateWorkers,
				MaxDrainRetries:      int(o.maxDrainRetries),
				EvictGracePeriod:     int(o.evictGracePeriod),
				WaitBetweenRotations: int(o.waitBetweenRotations),
				WaitBetweenDrains:    int(o.waitBetweenDrains),
			}

			rotatorMetadata, err := rotator.InitRotateCluster(&clusterRotator, &rotator.RotatorMetadata{}, logger)
			if err != nil {
				return err
			}
			if err = printJSON(rotatorMetadata); err != nil {
				return errors.Wrap(err, "failed to print cluster response")
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&whatamidoing, "wamid", "yourname", "testing")
	return cmd
}
