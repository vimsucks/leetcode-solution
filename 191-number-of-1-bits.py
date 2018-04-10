class Solution(object):
    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        i, count = 1, 0
        while i <= n:
            if n & i != 0:
                count += 1
            i = i << 1
        return count


if __name__ == "__main__":
    sln = Solution()
    print(sln.hammingWeight(123))
    print(sln.hammingWeight(1))
