import gradio as gr
from transformers import AutoModelForCausalLM, AutoTokenizer, pipeline
import torch
from langchain.document_loaders import GitLoader
from langchain.vectorstores import FAISS
from langchain.embeddings import HuggingFaceEmbeddings
from langchain.chains import RetrievalQA
from langchain.llms import HuggingFacePipeline

# ---- STEP 1: Load a Free AI Model (Mistral-7B) ---- #
MODEL_NAME = "mistralai/Mistral-7B-Instruct"
tokenizer = AutoTokenizer.from_pretrained(MODEL_NAME)
model = AutoModelForCausalLM.from_pretrained(
    MODEL_NAME, torch_dtype=torch.float16, device_map="auto"
)
generator = pipeline("text-generation", model=model, tokenizer=tokenizer, device=0)

# ---- STEP 2: Load Notes from Your GitHub Repo ---- #
REPO_URL = "https://github.com/nobitagit/knowledge.git"  # Change this!
loader = GitLoader(repo_path="repo", clone_url=REPO_URL)
documents = loader.load()

# ---- STEP 3: Create a Searchable Vector Database ---- #
vectorstore = FAISS.from_documents(documents, HuggingFaceEmbeddings())

# ---- STEP 4: Set Up AI Retrieval for Questions ---- #
llm = HuggingFacePipeline(pipeline=generator)
qa = RetrievalQA.from_chain_type(llm=llm, retriever=vectorstore.as_retriever())

# ---- STEP 5: Gradio Chat UI ---- #
def ask_question(query):
    return qa.run(query)

iface = gr.Interface(fn=ask_question, inputs="text", outputs="text", title="My Private AI")
iface.launch()
