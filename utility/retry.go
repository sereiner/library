package utility

//var CanStopErr = errors.New("retry can stop")

//func Retry(attempts int, sleep time.Duration, fn func() error) error {
//
//	if err := fn(); err != nil {
//		if errors.Is(err, CanStopErr) {
//			return err
//		}
//
//		if attempts--; attempts > 0 {
//			log.Printf("retry func error: %s. attemps #%d after %s.\n", err.Error(), attempts, sleep)
//			time.Sleep(sleep)
//			return Retry(attempts, 2*sleep, fn)
//		}
//		return err
//	}
//
//	return nil
//}
