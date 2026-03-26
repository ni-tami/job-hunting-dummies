import asyncio
import random
import string

from temporalio.client import Client
from temporalio.envconfig import ClientConfig
from temporalio.worker import Worker
from dataclasses import dataclass

from populate_company.activities.generate_company_activity import generate_company_activity

task_queue = "generate-company-py-queue"
workflow_name = "generate_company_py_workflow"


@dataclass(kw_only=True)
class CreateUser:
    username: str
    name: str

@dataclass(kw_only=True)
class CreateCompany:
    company_name: str
    user: CreateUser
    description: str
    website: str

async def main():
    # Create client to localhost on default namespace
    config = ClientConfig.load_client_connect_config()
    config.setdefault("target_host", "localhost:7233")
    client = await Client.connect(**config)

    # Run activity worker
    async with Worker(
        client, task_queue=task_queue, activities=[generate_company_activity]
    ):
        workflow_id = f'{workflow_name}-{"".join(random.choices(string.ascii_uppercase + string.digits, k=30))}'
        await client.execute_workflow(workflow_name, id=workflow_id, task_queue=task_queue)

if __name__ == "__main__":
    asyncio.run(main())
