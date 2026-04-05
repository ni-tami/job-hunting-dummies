import os

from llama_cpp import Llama


def main(example: str, model_path: str, filename: str, repo_id: str):
    """
    Sample chat response:
    {
        "id": "...",
        "object": "chat.completion",
        "created": 1775496089,
        "model": "...",
        "choices": [
            {
                "index": 0,
                "message": {
                    "role": "assistant",
                    "content": "Company AlphaTech specializes in providing cutting-edge technology solutions tailored to meet the high-performance demands of our customers. Leveraging state-of-the-art robotics platforms, we deliver unparalleled speed and cost efficiency for small-batch electronics production on-demand.\n\nOur platform excels at accelerating time-to-market processes by 10x while simultaneously reducing manufacturing costs significantly — a testament proven across countless industries including automotive innovation (e.g., Tesla), aerospace excellence applications like Zipline's projects in orbit; bolstering cutting-edge tech solutions underpins our customers' success, evidenced through the reliance of industry-leading clients such as NASA and CircuitHub.\n\nAlphaTech’s industrial-grade solution empowers manufacturers achieving peak performance while delivering unrivaled time-to-market capabilities — a shining beacon for competitors aiming to outpace rivals. Our proprietary platforms stand unyielding in powering competitive advantages proven across industries like Aerospace & Automotive Industries’ innovators, bolstering unparalleled excellence benchmark demonstrated by clients' reliance on industry-leading stalwarts such as Tesla and CircuitHub.\n\nAlphaTech’s cutting-edge tech solutions tailored effortlessly seamlessly streamline your manufacturing processes – a stellar solution proving our prowess. Cutting edge technological platforms empower manufacturers achieving peak performance while consistently outperforming rivals — an exceptional competitor with proven time-to-market capabilities rivalled through competitors’ alliance to sur",
                },
                "logprobs": None,
                "finish_reason": "length",
            }
        ],
        "usage": {"prompt_tokens": 261, "completion_tokens": 251, "total_tokens": 512},
    }

    """
    llm = Llama.from_pretrained(
        repo_id=repo_id,
        local_dir=model_path,
        filename=filename,
    )

    resp = llm.create_chat_completion(
        max_tokens=640,
        repeat_penalty=1.4,
        temperature=0.8,
        messages=[
            {
                "role": "system",
                "content": "You are a copywriter who creates concise, professional company descriptions for fictional businesses.",
            },
            {
                "role": "user",
                "content": ("Write a description for a fictional company based on the style and structure of the following real company description. "
                            "Follow these rules:\n"
                            "1. Create at most 6 sentences overall.\n"
                            "2. Keep the same industry category as the original example.\n"
                            "3. Do not copy any exact phrases, entity names, or proper nouns from the original example.\n"
                            "4. Invent a new company name, what the company does, a named internal platform or product, and fictional customer names.\n"
                            "5. Highlight key differentiator or technology in one sentence, and/or mention the customer benefits (e.g. speed, cost, time-to-market) in one or two sentences.\n"
                            "6. Do not explain platform or product longer than 2 sentences.\n"
                            "7. Keep the tone B2B-professional.\n"
                            "8. Do not reference the example original company in the output.\n"
                            "9. Do not explain products.\n"
                            "Original example company description:\n"
                            f"\"\"\"\n{example}\n\"\"\""
                            ),
            },
        ]
    )
    print(resp["choices"][0]["message"]["content"])


if __name__ == "__main__":
    example = os.environ.get("example")
    model_path = os.environ.get("model_path")
    filename = os.environ.get("filename")
    repo_id = os.environ.get("repo_id")
    if example and model_path and filename:
        print(f"Regenerate description from example: {example}.")
        main(example, model_path, filename, repo_id)
        print("Done.")
    else:
        print("Recheck params -_")
