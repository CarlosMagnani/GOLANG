package workerpools

import "fmt"

//WORKER POOLS SÃO COMO UMA FILA DE TAREFAS E ASSIM PODEMOS TER VARIOS PROCESSOS PARA FAZER OS ITENS DA FILA DISTRIBUINDO A CARGA
// E AGILIZANDO AINDA MAIS O PROCESSO

//FUNCS RECURSIVAS DEPENDEM DELAS MESMAS PARA CONTINUAR SENDO EXECUTADAS, CHAMAM ELAS MESMAS

//É NECESSÁRIO MOSTRAR PARA O GO QUANDO PARAR, POR CONTA DO 'STACKEROVERFLOW' ONDE SOBRECARREGA O SISTEMA
func fibonnaci(posicao int) int  {
	//CONDIÇÃO DE PARADA
	if posicao <= 1{
		return posicao
	}

	return fibonnaci(posicao - 2) + fibonnaci(posicao - 1)
}


func main()  {

	// DOIS CANAIS PARA ARMAZENAR OS ITENS DE TAREFAS E SEUS RESULTADOS 
	tarefas := make(chan int, 100)
	resultados := make(chan int, 100)

	// AQUI INICIALIZO MINHA ROTINA DE WORKER PASSANDO MEUS DOIS CANAIS CRIADOS 
	// COM OS MEUS WORKERS EU POSSO CHAMA-LOS MAIS DE UMA VEZ FAZENDO ASSIM A DIVISÃO DA CARGA DE PROCESSOS
	// E PUXANDO MAIS TAREFAS DA MINHA FILA
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)





	//AQUI EU POSSO MANIPULAR QUANTOS NUMEROS EU QUERO QUE MINHA FUNÇÃO DE FIBONNACI TRAGA 
	// SEGUINDO O PADRÃO DE WORKER POOLS EU UTILIZO MEU CANAL DE TAREFAS PARA PODER REGISTRAR A QUANTIDADE DE NUMEROS
	//E POSTERIORMENTE ESSES NUMEROS VÃO RESOLVIDOS PELA MINHA FUNÇÃO DE FIBONNACI UTILIZANDO MEUS WORKERS
	for i := 0; i < 80; i++ {
		tarefas <- i
	}
	// AQUI EU POSSO FECHAR MEU CANAL APÓS EU JÁ TER POPULADO MEU CANAL COM OS DADOS QUE EU PRECISO QUE MINHA FUNÇÃO PRECISE FAZER 
	close(tarefas)	

	//AQUI EU CONSIGO VISUALIZAR A REALIZAÇÃO DAS MINHAS TAREFAS ATRAVÉS DO MEU CANAL DE RESULTADOS
	for i := 0; i < 80; i++ {
		resultado := <- resultados
		fmt.Println(resultado)
	}
}	
//A UTILIDADE PRIMORDIAL DE SE UTILIZAR OS WORKERS É QUE CONFORME VAI CHEGANDO "TAREFAS" PELO CANAL ELAS JÁ VÃO SENDO REALIZADAS GERANDO MAIS PERFOMANCE PARA NOSSO PROJETO

//AO UTILIZAR A NOMECLATURA DE '<-CHAN INT' ESTOU DIZENDO QUE ESTE CANAL É UM CANAL QUE SOMENTE RECEBE DADOS
//JÁ O CONTRARIO ONDE TEMOS 'CHAN<-INT' ESTAMOS DIZENDO QUE É UM CANAL SOMENTE DE ENVIO DE DADOS 
func worker(jobs <-chan int, results chan<- int)  {
	// AQUI EU ESTOU ITERANDO EM TODOS OS MEUS JOBS E ENVIANDO CADA JOB PARA MINHA FUNÇÃO DE FIBONNACI E PASSANDO O RETORNO DELA PARA O RESULTADO
	for job := range jobs {
		results <- fibonnaci(job)
	}
}