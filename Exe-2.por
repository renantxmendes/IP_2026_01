programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() 
  {
  inteiro quant_jogos, x
  
  escreva("Número de jogos: ")
  leia(quant_jogos)

    real total_pessoas, pp, pg, pa, pc
    escreva("Qual total de incressos vendidos? ")
    leia(total_pessoas)

    escreva("Quantos ingressos da categoria popular foram vendidos? ")
    leia(pp)

    escreva("Quantos ingressos da categoria geral foram vendidos? ")
    leia(pg)

    escreva("Quantos ingressos da categoria arquibancada foram vendidos? ")
    leia(pa)

    escreva("Quantos ingressos da categoria cadeiras foram vendidos? ")
    leia(pc)

    limpa()

    real renda = 
    total_pessoas * pp / 100 * 1 +
    total_pessoas * pg / 100 * 5 + 
    total_pessoas * pa / 100 * 10 + 
    total_pessoas * pc / 100 * 20

    escreva("A renda do JOGO N.", x , " é = ", "R$ ", renda, "\n")
   }
  }
