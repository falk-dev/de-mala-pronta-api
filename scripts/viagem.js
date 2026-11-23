const API = "/api/viagens";
const campos = ["titulo", "data_inicio", "data_fim", "orcamento", "status"];

function el(id) {
  return document.getElementById(id);
}

function mostrarMensagem(texto) {
  el("mensagem").textContent = texto;
}

function abrirModal(titulo) {
  el("modal-titulo").textContent = titulo;
  el("erro").textContent = "";
  el("modal").showModal();
}

function novaViagem() {
  el("form").reset();
  el("id").value = "";
  abrirModal("Nova viagem");
}

function fecharModal() {
  el("modal").close();
}

async function listar() {
  const resposta = await fetch(API);
  const viagens = await resposta.json();
  const tabela = el("tabela");
  tabela.innerHTML = "";

  for (const v of viagens) {
    const totalResp = await fetch(`${API}/${v.id}/total-despesas`);
    const total = (await totalResp.json()).total;

    const tr = document.createElement("tr");
    for (const valor of [v.id, v.titulo, v.data_inicio, v.data_fim, v.orcamento, v.status, total]) {
      const td = document.createElement("td");
      td.textContent = valor;
      tr.appendChild(td);
    }

    const acoes = document.createElement("td");
    const editar = document.createElement("button");
    editar.textContent = "Editar";
    editar.onclick = () => editarViagem(v);
    const excluir = document.createElement("button");
    excluir.textContent = "Excluir";
    excluir.onclick = () => excluirViagem(v.id);
    acoes.append(editar, excluir);
    tr.appendChild(acoes);

    tabela.appendChild(tr);
  }
}

function editarViagem(v) {
  el("id").value = v.id;
  for (const c of campos) {
    el(c).value = v[c];
  }
  abrirModal("Editar viagem");
}

async function salvar(evento) {
  evento.preventDefault();
  const id = el("id").value;
  const dados = {
    titulo: el("titulo").value,
    data_inicio: el("data_inicio").value,
    data_fim: el("data_fim").value,
    orcamento: parseFloat(el("orcamento").value),
    status: el("status").value,
  };
  if (isNaN(dados.orcamento)) {
    dados.orcamento = 0;
  }

  const resposta = await fetch(id ? `${API}/${id}` : API, {
    method: id ? "PUT" : "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(dados),
  });

  if (resposta.ok) {
    mostrarMensagem("Viagem salva com sucesso.");
    fecharModal();
    listar();
  } else {
    el("erro").textContent = (await resposta.json()).erro;
  }
}

async function excluirViagem(id) {
  if (!confirm("Excluir esta viagem?")) {
    return;
  }
  const resposta = await fetch(`${API}/${id}`, { method: "DELETE" });
  if (resposta.ok) {
    mostrarMensagem("Viagem excluída.");
    listar();
  } else {
    mostrarMensagem((await resposta.json()).erro);
  }
}

el("form").addEventListener("submit", salvar);
el("cancelar").addEventListener("click", fecharModal);
el("novo").addEventListener("click", novaViagem);
listar();
