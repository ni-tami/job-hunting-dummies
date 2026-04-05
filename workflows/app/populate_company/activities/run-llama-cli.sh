llama-cli -m secret/model.gguf \
  -c 640 \
  --temp 0.8 \
  --repeat-penalty 1.4 \
  -sys "You are a copywriter who creates concise, professional company descriptions for fictional businesses." \
  -p "Write a description for a fictional company based on the style and structure of the following real company description. Follow these rules:\n1. Create at most 6 sentences overall.\n2. Keep the same industry category as the original.\n3. Do not copy any exact phrases, entity names, or proper nouns from the original.\n4. Invent a new company name, a named internal platform or product (mention only, not explained), and fictional (but plausible) customer names.\n6. Mirror the structure: state what the company does, highlight a key differentiator or technology in one sentence, and mention the customer benefit in one or two sentences (e.g. speed, cost, time-to-market).\n5. Keep the tone confident and B2B-professional.\n7. Do not mention or reference the original company in the output.\n8.Do not list and explain products.\nOriginal company description:\n\"\"\"\n[INSERT HERE]\n\"\"\""
