# import sys

# MAX = sys.maxsize

def coin_change_n_ways(Coins, m, total):
    # Instatiate memo table for tablization

    memo = [0] * (total + 1)

    # There is 1 way to get zero, no coins
    memo[0] = 1

    for coin in Coins:
        for j in range(coin, total + 1):
            memo[j] += memo[j - coin]
    
    # print(f'{memo}')
    return memo[total]

def coin_change_min_ways(Coins, m, amount):
    MAX = float('inf') # a number bigger than all others
    dp = [0] + [MAX] * amount

    for i in range(1, amount + 1):
        dp[i] = min([dp[i - c] if i - c >= 0 else MAX for c in Coins]) + 1

    return [dp[amount], -1][dp[amount] == MAX]


if __name__ == '__main__':
    # Python3 f strings
    print(f'{coin_change_n_ways([1, 2, 5, 10], 4, 5)} ways to make 5')
    print(f'{coin_change_n_ways([1, 2, 5, 10], 4, 99)} ways to make 99')

    print(
        f'{coin_change_min_ways([1, 2, 4, 10], 4, 5)} min # of ways to make 5')
    print(
        f'{coin_change_min_ways([1, 2, 5, 10], 4, 99)} min # of ways to make 99')
