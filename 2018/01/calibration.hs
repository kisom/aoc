import Data.Set (insert, member, singleton)

calibrate :: [Int] -> Int -> Int
calibrate deltas frequency = sum(deltas) + frequency

detectRepeated :: [Int] -> Int -> Int
detectRepeated deltas frequency = detector (singleton frequency) deltas frequency
  where 
    detector seen [] frequency = detector seen deltas frequency
    detector seen (delta:_deltas) frequency = if member frequency_ seen
       then frequency_
        else detector (insert frequency_ seen) _deltas frequency_
        where frequency_ = frequency + delta

-- Utility functions

processDelta :: String -> Int
processDelta s@(x:xs) = if x == '+' then read xs else read s

getDeltas :: [String] -> [Int]
getDeltas = map processDelta 

main :: IO ()
main = do 
  deltas <- (readFile "calibration.txt" >>= return . lines >>= return . getDeltas)
  let calibrated = show $ calibrate deltas 0
  putStrLn ("Calibrated: " ++ calibrated)

  let repeated = show $ detectRepeated deltas 0
  putStrLn ("Repeated: " ++ repeated)