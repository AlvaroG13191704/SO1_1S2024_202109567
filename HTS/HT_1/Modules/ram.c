#include <linux/module.h>
// para usar KERN_INFO
#include <linux/kernel.h>
// para info_ram
#include <linux/mm.h>


//Header para los macros module_init y module_exit
#include <linux/init.h>
//Header necesario porque se usara proc_fs
#include <linux/proc_fs.h>
/* for copy_from_user */
#include <asm/uaccess.h>
/* Header para usar la lib seq_file y manejar el archivo en /proc*/
#include <linux/seq_file.h>

const long gigabyte = 1024 * 1024 * 1024;

// Obteniendo estadísticas del sistema
struct sysinfo si;

static void init_meminfo(void) {
   si_meminfo(&si);
}


MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de RAM, Laboratorio Sistemas Operativos 1");
MODULE_AUTHOR("AlvaroG");

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int escribir_archivo(struct seq_file *archivo, void *v)
{
    unsigned long totalram, freeram, usedram, percent_used;

    init_meminfo();

    // Extract free memory
    freeram = si.freeram / gigabyte;
    // Extract total memory
    totalram = si.totalram / gigabyte;

    // Calculate used memory
    usedram = (totalram - freeram) / gigabyte;

    // Calculate percentage of used memory
    percent_used = usedram * 100 / totalram;

    // Print all the information to the file
    seq_printf(archivo, "Free memory: %lu\n", freeram);
    seq_printf(archivo, "Total memory: %lu\n", totalram);
    seq_printf(archivo, "Used memory: %lu\n", usedram);
    seq_printf(archivo, "Percentage of memory used: %lu\n", percent_used);

    return 0;
}

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

//Funcion a ejecuta al insertar el modulo en el kernel con insmod
static int _insert(void)
{
    proc_create("ram_202109567", 0, NULL, &operaciones);
    printk(KERN_INFO "202109567\n");
    return 0;
}

//Funcion a ejecuta al remover el modulo del kernel con rmmod
static void _remove(void)
{
    remove_proc_entry("ram_202109567", NULL);
    printk(KERN_INFO "Laboratorio Sistemas Operativos 1\n");
}

module_init(_insert);
module_exit(_remove);