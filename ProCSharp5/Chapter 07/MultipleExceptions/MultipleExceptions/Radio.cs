﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace MultipleExceptions
{
    class Radio
    {
        public void TurnOn(bool on)
        {
            if(on)
                Console.WriteLine("Radio edit.");
            else
                Console.WriteLine("Radio silence.");
        }
    }
}