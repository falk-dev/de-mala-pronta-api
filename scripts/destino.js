const API = "/api/destinos";
const campos = ["viagem_id", "cidade", "pais", "data_chegada", "data_saida"];

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

function novoDestino() {
  el("form").reset();
  el("id").value = "";
  abrirModal("Novo destino");
}

function fecharModal() {
  el("modal").close();
}

const titulos = {};

async function carregarViagens() {
  const resposta = await fetch("/api/viagens");
  const viagens = await resposta.json();
  const select = el("viagem_id");
  select.innerHTML = "";
  for (const v of viagens) {
    const opcao = document.createElement("option");
    opcao.value = v.id;
    opcao.textContent = v.titulo;
    titulos[v.id] = v.titulo;
    select.appendChild(opcao);
  }
}

async function listar() {
  const resposta = await fetch(API);
  const destinos = await resposta.json();
  const tabela = el("tabela");
  tabela.innerHTML = "";

  for (const d of destinos) {
    const tr = document.createElement("tr");
    for (const valor of [d.id, titulos[d.viagem_id], d.cidade, d.pais, d.data_chegada, d.data_saida]) {
      const td = document.createElement("td");
      td.textContent = valor;
      tr.appendChild(td);
    }

    const acoes = document.createElement("td");
    const editar = document.createElement("button");
    editar.textContent = "Editar";
    editar.onclick = () => editarDestino(d);
    const excluir = document.createElement("button");
    excluir.textContent = "Excluir";
    excluir.onclick = () => excluirDestino(d.id);
    acoes.append(editar, excluir);
    tr.appendChild(acoes);

    tabela.appendChild(tr);
  }
}

function editarDestino(d) {
  el("id").value = d.id;
  for (const c of campos) {
    el(c).value = d[c];
  }
  abrirModal("Editar destino");
}

async function salvar(evento) {
  evento.preventDefault();
  const id = el("id").value;
  const dados = {
    viagem_id: parseInt(el("viagem_id").value) || 0,
    cidade: el("cidade").value,
    pais: el("pais").value,
    data_chegada: el("data_chegada").value,
    data_saida: el("data_saida").value,
  };

  const resposta = await fetch(id ? `${API}/${id}` : API, {
    method: id ? "PUT" : "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(dados),
  });

  if (resposta.ok) {
    mostrarMensagem("Destino salvo com sucesso.");
    fecharModal();
    listar();
  } else {
    el("erro").textContent = (await resposta.json()).erro;
  }
}

async function excluirDestino(id) {
  if (!confirm("Excluir este destino?")) {
    return;
  }
  const resposta = await fetch(`${API}/${id}`, { method: "DELETE" });
  if (resposta.ok) {
    mostrarMensagem("Destino excluído.");
    listar();
  } else {
    mostrarMensagem((await resposta.json()).erro);
  }
}

el("form").addEventListener("submit", salvar);
el("cancelar").addEventListener("click", fecharModal);
el("novo").addEventListener("click", novoDestino);
carregarViagens().then(listar);
