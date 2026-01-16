data Player = X | O | N
  deriving (Eq)

instance Show Player where
    show X = "X"
    show O = "O"
    show N = "_" 
    
type Board = [Player]

displayBoard :: Board -> String
displayBoard [] = []
displayBoard (p1:p2:p3:ps) = show p1 ++ show p2 ++ show p3 ++ "\n" ++ displayBoard ps

rowWinner :: Player -> Board -> Bool
rowWinner _ [] = False
rowWinner p (p1:p2:p3:ps)
 | (p == p1) && (p == p2) && (p == p3) = True
 | otherwise = rowWinner p ps

transpose :: Board -> Board
transpose [p1, p2, p3, p4, p5, p6, p7, p8, p9] = [p1, p4, p7, p2, p5, p8, p3, p6, p9]

diagWinner :: Player -> Board -> Bool
diagWinner p [p1,p2,p3,p4,p5,p6,p7,p8,p9] 
 | (p == p1) && (p == p5) && (p == p9) = True
 | (p == p3) && (p == p5) && (p == p7) = True
 | otherwise = False

winner :: Player -> Board -> Bool 
winner p b
  | rowWinner p b = True
  | rowWinner p (transpose b) = True
  | diagWinner p b = True
  | otherwise = False
 
updateBoard :: Int -> Player -> Board -> Board
updateBoard 0 player (b:bs) = [player] ++ bs
updateBoard idx player (b:bs) = [b] ++ updateBoard (idx - 1) player bs

gameLoop :: [Player] -> Board -> IO ()
gameLoop [] _ = print "Game over!"
gameLoop (p:ps) b = do 
  putStrLn (displayBoard b)
  putStrLn ("It's player " ++ show p ++ "'s turn")
  putStrLn "Enter a Move Index (1-9)"
  move <- getLine
  let moveInt = read move :: Int
  let b2 = updateBoard (moveInt - 1) p b
  if winner p b2 then do
    putStrLn (displayBoard b2)
    putStrLn ("Player " ++ show p ++ " wins")
  else 
    gameLoop ps b2

main :: IO ()
main = do
  putStrLn "Welcome to Tic Tac Toe!"
  let turns = [O,X,O,X,O,X,O,X,O]
  let board = [N,N,N,N,N,N,N,N,N]
  gameLoop turns board

-- [O,O,O,O,O,O,O,O,O]