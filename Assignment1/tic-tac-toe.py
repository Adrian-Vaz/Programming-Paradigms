winning_combo = [[0,1,2], [4,5,6], [8,9,10], [0,5,6], [1,5,9], [2,6,10], [3,7,11], [2,4,6], [0,5,10], [1,6,11], [2,5,8], [3,6,9]]
board = ["-"] * 12
move_history = []

def display_board():
    print(board[0], board[1], board[2], board[3])
    print(board[4], board[5], board[6], board[7])
    print(board[8], board[9], board[10], board[11])


def winner(player):
    for combo in winning_combo:
        if board[combo[0]] == player and board[combo[1]] == player and board[combo[2]] == player:
            return True
    return False
    
def take_turn(player):
    print("It's player " + player + "'s turn")
    move = int(input("Enter move index (1-12): "))
    board[move-1] = player
    move_history.append((move-1, player))

def remove_move():
    if len(move_history) > 0:
        position, player = move_history.pop()
        board[position] = "-"
        print("Move removed! Last move by player " + player + " at position " + str(position+1) + " has been undone.")
    else:
        print("No moves to remove.")
    
def game_loop():
    for player in ["o","x","o","x","o","x","o","x","o","x","o","x"]:
        display_board()
        if len(move_history) % 3 == 0 and len(move_history) > 0:
            remove_choice = input("Remove last move? (y/n): ").lower()
            if remove_choice == "y":
                remove_move()
                continue
        take_turn(player)
        if winner(player):
            display_board()
            print("Player " + player + " wins!")
            break
def main():
    game_loop()

if __name__ == "__main__":
    main()