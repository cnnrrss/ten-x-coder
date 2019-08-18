# For example, here's several helpful packages to load in 

import numpy as np # linear algebra
import pandas as pd # data processing, CSV file I/O (e.g. pd.read_csv)

# Input data files are available in the "../input/" directory.
# For example, running this (by clicking run or pressing Shift+Enter) will list all files under the input directory

import os
for dirname, _, filenames in os.walk('/kaggle/input'):
    file_to_open = ""
    df = pd.DataFrame()
    if len(filenames) > 0:
        file_to_open = filenames[0]
    if file_to_open != "":
        df = pd.read_fwf(os.path.join(dirname, file_to_open))
        print(df.columns)
        print(df.iat[0,0])
#         print(df[1])
# Any results you write to the current directory are saved as output.
