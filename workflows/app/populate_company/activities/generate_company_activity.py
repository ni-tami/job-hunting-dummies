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
async def generate_company_activity() -> CreateCompany:
    create_company_payload = CreateCompany(
        company_name="FamInc",
        user=CreateUser(username="admin", name="admin andrea"),
        description="Unc Inc.",
        website="https://www.unc.inc",
    )
    activity.logger.info(
        "Generate activity from python activity:", create_company_payload
    )
    return create_company_payload
