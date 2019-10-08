# Amazon QuickSight

- To join tables both are based on the same SQL database data source.
- To join tables from _different_ data sources, create the join **before** importin to Amazon Quicksight (Quicksight does not provide facility to join tables from different data sources)
- QuickSight allows configure join type (inner, outer, left, right)
- If you choose a table and made changes to the fields (ex: changing field name), these changes are discarded when you add tables using the join interface

### Chart Types

**Tabular Reports** - customized table view of your data.

**Heat Maps** - use to show a measure for the **intersection of two dimensions**, with color-coding to easily differentiate where values fall in the range. Heat maps can also be used to show the count of values for the intersection of two dimensions.

**Line Chart** - used to compare changes in measure values over a period of time for following scenarios:
- One measure over a period of time (gross sales by month)
- Multiple measures over a period of time (gross sales and net sales by month)
- One measure for a dimension over a period of time (# of flight delays per day by airline)

**Tree Maps** - used to visualize one or two measures for a dimension