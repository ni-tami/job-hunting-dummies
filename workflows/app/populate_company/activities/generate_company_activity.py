import asyncio
import random
import string

from temporalio import activity
from temporalio.client import Client
from temporalio.envconfig import ClientConfig
from temporalio.worker import Worker
from dataclasses import dataclass

task_queue = "generate-company-py-queue"
workflow_name = "generate-company-workflow"
activity_name = "generate-company-activity"


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

@activity.defn(name=activity_name)
async def generate_company_activity() -> CreateCompany:
    create_company_payload = CreateCompany(
        company_name="FamInc",
        user=CreateUser(
            username="admin",
            name="admin andrea"
        ),
        description="Unc Inc.",
        website="https://www.unc.inc",
    )
    activity.logger.info(
        "Generate activity from python activity:", create_company_payload
    )
    return create_company_payload

# async def main():
#     # Create client to localhost on default namespace
#     config = ClientConfig.load_client_connect_config()
#     config.setdefault("target_host", "localhost:7233")
#     client = await Client.connect(**config)

#     # Run activity worker
#     async with Worker(client, task_queue=task_queue, activities=[generate_company_activity]) as worker:
#         # Run the Go workflow
#         await worker.run()
#         # result = await client.execute_workflow(
#         #     workflow_name, "PY Temporal", id=workflow_id, task_queue=task_queue
#         # )
#         # # Print out "Hello, PY Temporal!"
#         # print(result)


# if __name__ == "__main__":
#     asyncio.run(main())
