class Solution {
    public int[][] flipAndInvertImage(int[][] A) {
        if (A==null || A.length==0) return null;
        for (int i=0; i<A.length; i++) 
	    {
			// check out the j, k declarations
			for(int j=0, k=A[i].length-1; j<=k; j++, k--) 
			{
				int temp1 = (A[i][j]==0) ? 1 : 0;
				int temp2 = (A[i][k]==0) ? 1 : 0;
				A[i][j] = temp2;
				A[i][k] = temp1;
			}
	    }
    
        return A;
    }
}