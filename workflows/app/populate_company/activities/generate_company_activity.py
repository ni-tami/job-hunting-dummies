from typing import List

from temporalio import activity
from dataclasses import dataclass

activity_name = "generate-company-py-activity"


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
async def generate_company_activity(sources: List[CreateCompany]) -> List[CreateCompany]:
    create_companies_payload = []
    for source in sources:
        create_companies_payload.append(
            CreateCompany(
                company_name="hehe "+source.company_name,
                user=CreateUser(username=source.user.username, name=source.user.name),
                description="hehe "+source.description,
                website=source.website,
            )
        )
    activity.logger.info(
        f"Generate {len(create_companies_payload)} activities from python activity."
    )
    return create_companies_payload
