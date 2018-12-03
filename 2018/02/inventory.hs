import Control.Applicative
import Data.List (group, sort)

-- lots of tiny, testable pieces
countElements :: String -> [Int]
countElements word = map (\x -> length x) . group $ sort word

folder :: Int -> (Int, Int) -> (Int, Int)
folder c (two, three)
  | c == 2 = (min (two + 1) 1, three) 
  | c == 3 = (two, min (three + 1) 1) 
  | otherwise = (two, three)

addPair :: (Int, Int) -> (Int, Int) -> (Int, Int)
addPair (a, b) (x, y) = (a+x, b+y)

reduceWordCount :: [Int] -> (Int, Int)
reduceWordCount = foldr folder (0, 0)

buildInventory :: [String] -> [(Int, Int)]
buildInventory = map reduceWordCount . map countElements

countInventory :: [(Int, Int)] -> (Int, Int)
countInventory = foldr addPair (0, 0)

checksum :: (Int, Int) -> Int
checksum (x, y) = x * y

compareWords :: String -> String -> (String, Int) -> (String, Int)
compareWords [] [] (common, differences) = (common, differences)
compareWords (x:xs) (y:ys) (common, differences)
  | differences > 1 = (common, differences)
  | x == y = compareWords xs ys (common ++ [x], differences)
  | otherwise = compareWords xs ys (common, differences + 1)

countDifferences :: String -> String -> (String, Int)
countDifferences x y = compareWords x y ("", 0)

doesMatch :: (String, Int) -> Bool
doesMatch = (==) 1 . snd

findMatchFor :: String -> [String] -> Maybe String
findMatchFor x [] = Nothing
findMatchFor x (y:ys) 
  | doesMatch (countDifferences x y) = Just . fst $ countDifferences x y
  | otherwise = findMatchFor x ys

findMatch :: [String] -> Maybe String
findMatch [] = Nothing
findMatch (x:[]) = Nothing
findMatch (x:xs) = findMatchFor x xs <|> findMatch xs

showMatch :: Maybe String -> String
showMatch (Just s) = s
showMatch Nothing  = "no matches found"

main :: IO ()
main = do
    boxes <- (readFile "inventory.txt" >>= return . lines)
    let cksum = checksum . countInventory $ buildInventory boxes
    putStrLn $ "Checksum: " ++ (show cksum)
    putStrLn $ "Common: " ++ (showMatch $ findMatch boxes)