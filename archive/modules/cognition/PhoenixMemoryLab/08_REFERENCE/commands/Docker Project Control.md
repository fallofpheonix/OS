# Docker Project Control

## Stop All Projects

```bash
docker stop $(docker ps -q)
```

## Start One Project

```bash
cd ~/engineering/workspace/active/project-name
docker compose up -d
```

## Check Running Ports

```bash
docker ps --format "table {{.Names}}\t{{.Ports}}"
```

## Standard Project Makefile

```makefile
run:
	docker compose up -d

stop:
	docker compose down

test:
	pytest tests/

lint:
	ruff check .

build:
	docker compose build

clean:
	docker compose down -v
	find . -type d -name __pycache__ -exec rm -rf {} +

logs:
	docker compose logs -f
```

