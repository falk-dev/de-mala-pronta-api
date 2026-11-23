const API = "/api/despesas";
const campos = ["viagem_id", "descricao", "valor", "data", "categoria"];

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

function novaDespesa() {
  el("form").reset();
  el("id").value = "";
  abrirModal("Nova despesa");
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
  const despesas = await resposta.json();
  const tabela = el("tabela");
  tabela.innerHTML = "";

  for (const d of despesas) {
    const tr = document.createElement("tr");
    for (const valor of [d.id, titulos[d.viagem_id], d.descricao, d.valor, d.data, d.categoria]) {
      const td = document.createElement("td");
      td.textContent = valor;
      tr.appendChild(td);
    }

    const acoes = document.createElement("td");
    const editar = document.createElement("button");
    editar.textContent = "Editar";
    editar.onclick = () => editarDespesa(d);
    const excluir = document.createElement("button");
    excluir.textContent = "Excluir";
    excluir.onclick = () => excluirDespesa(d.id);
    acoes.append(editar, excluir);
    tr.appendChild(acoes);

    tabela.appendChild(tr);
  }
}

function editarDespesa(d) {
  el("id").value = d.id;
  for (const c of campos) {
    el(c).value = d[c];
  }
  abrirModal("Editar despesa");
}

async function salvar(evento) {
  evento.preventDefault();
  const id = el("id").value;
  const dados = {
    viagem_id: parseInt(el("viagem_id").value) || 0,
    descricao: el("descricao").value,
    valor: parseFloat(el("valor").value) || 0,
    data: el("data").value,
    categoria: el("categoria").value,
  };

  const resposta = await fetch(id ? `${API}/${id}` : API, {
    method: id ? "PUT" : "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(dados),
  });

  if (resposta.ok) {
    mostrarMensagem("Despesa salva com sucesso.");
    fecharModal();
    listar();
  } else {
    el("erro").textContent = (await resposta.json()).erro;
  }
}

async function excluirDespesa(id) {
  if (!confirm("Excluir esta despesa?")) {
    return;
  }
  const resposta = await fetch(`${API}/${id}`, { method: "DELETE" });
  if (resposta.ok) {
    mostrarMensagem("Despesa excluída.");
    listar();
  } else {
    mostrarMensagem((await resposta.json()).erro);
  }
}

el("form").addEventListener("submit", salvar);
el("cancelar").addEventListener("click", fecharModal);
el("novo").addEventListener("click", novaDespesa);
carregarViagens().then(listar);
