
					Manuel Utilisateur
•	Compiler votre programme à debugger en static avec l’option débogage (-g)

•	Compiler le tout avec : go build

•	Lancer cette commande : export GODEBUG=cgocheck=0

•	Lancer le programme avec en indiquant le (chemin) fichier source du programme à debugger en premier suivi  (chemin) de son code compilé: 
							Example :  ./Ressources  essai.c  essai
															Ressources : programme compilé
															essai.c : code source à debugger
															essai : code compilé d’ssai
 

•	Choisir un breakpoint ou plusieurs en cliquant sur un des numéros de ligne 

•	Appuyer sur le bouton Play pour commencer le débogage

•	Appuyer sur le bouton avance pour faire un débogage pas à pas 

•	 Appuyer sur le bouton retour pour revenir sur vos pas

•	Appuyer sur le bouton avance rapide pour faire un saut 

•	Appuyer sur le bouton rembobiner pour retourner du saut

•	Les frames et variables par frame sont visibles dans rectangle de gauche
