# job-hunting-dummies
A job search app for dummy HR and dummy job hunters. This is planned to be a sim game fingers crossed.

## How to Run
Before mi forget how to run

### Run temporal server and local cockroach database
Pull temporal repo and compose
```bash
git clone https://github.com/temporalio/docker-compose.git
cd docker-compose
podman compose up -f docker-compose-postgres.yml -d
```

Pull cockroach image, set volume mount paths and initialize database.

### Run backend service
Make sure [.env](./service/.env) is set up.  
From project repository root, run:
```bash
cd service
source .env
make runserver
```

### (Optional) Run temporal python activity handler
Python activity handler is used in populate company usecase so run worker if needed for this workflow.  
Make sure environment is synced: `uv sync`  
From project repository root, run:
```bash
cd workflows/app
make runpyworker wf=populate_company
```

### Run temporal go workflow and activity handler
Go activity and workflow handler is used in:
* populate_user
* populate_company (Run python activity handler first)

Make sure dependencies are added: `go get`  
From project repository root, run:
```bash
cd workflows/app
make runworker wf=populate_company
# or make runworker wf=populate_user
```

### Run temporal go workflow starter
Go workflow is initiated in paused state. Trigger workflow manually from temporal web UI or run workflow starter. Go starter is used in:  
* populate_user
* populate_company

Make sure dependencies are added: `go get`  
From project repository root, run:
```bash
cd workflows/app
make runstarter wf=populate_company
# or make runstarter wf=populate_user
```

## Entities
[DB Diagram file](./service/docs/diagram)
