from sqlalchemy import create_engine, text

engine = create_engine("sqlite+pysqlite:///:memory:", echo=True)
with engine.connect() as conn:
    conn.execute(text("CREATE TABLE some_table (x int, y int)"))
    conn.execute(
        text("INSERT INTO some_table (x, y) VALUES (:x, :y)"),
        [{"x": 1, "y": 1}, {"x": 2, "y": 4}, {"x": 6, "y": 8}, {"x": 9, "y": 10}],
    )
    conn.commit()

# with engine.connect() as conn:
#     result = conn.execute(text("SELECT x, y FROM some_table"))
#     for row in result:
#         print(f"x: {row.x}  y: {row.y}")

with engine.connect() as conn:
    result = conn.execute(text("SELECT x, y FROM some_table"))
    for y, x in result:
        print(f"x: {x}  y: {y}")
"""
❯ python fetching_rows.py
2023-03-05 17:50:50,499 INFO sqlalchemy.engine.Engine BEGIN (implicit)
2023-03-05 17:50:50,499 INFO sqlalchemy.engine.Engine CREATE TABLE some_table (x int, y int)
2023-03-05 17:50:50,499 INFO sqlalchemy.engine.Engine [generated in 0.00019s] ()
2023-03-05 17:50:50,500 INFO sqlalchemy.engine.Engine INSERT INTO some_table (x, y) VALUES (?, ?)
2023-03-05 17:50:50,500 INFO sqlalchemy.engine.Engine [generated in 0.00011s] [(1, 1), (2, 4), (6, 8), (9, 10)]
2023-03-05 17:50:50,500 INFO sqlalchemy.engine.Engine COMMIT
2023-03-05 17:50:50,500 INFO sqlalchemy.engine.Engine BEGIN (implicit)
2023-03-05 17:50:50,500 INFO sqlalchemy.engine.Engine SELECT x, y FROM some_table
2023-03-05 17:50:50,500 INFO sqlalchemy.engine.Engine [generated in 0.00012s] ()
x: 1  y: 1
x: 2  y: 4
x: 6  y: 8
x: 9  y: 10
2023-03-05 17:50:50,500 INFO sqlalchemy.engine.Engine ROLLBACK
"""