/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.cbus.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public abstract class ParameterValue implements Message {

  // Abstract accessors for discriminator values.
  public abstract ParameterType getParameterType();

  // Arguments.
  protected final Short numBytes;

  public ParameterValue(Short numBytes) {
    super();
    this.numBytes = numBytes;
  }

  protected abstract void serializeParameterValueChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ParameterValue");

    // Switch field (Serialize the sub-type)
    serializeParameterValueChild(writeBuffer);

    writeBuffer.popContext("ParameterValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ParameterValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static ParameterValue staticParse(
      ReadBuffer readBuffer, ParameterType parameterType, Short numBytes) throws ParseException {
    readBuffer.pullContext("ParameterValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    ParameterValueBuilder builder = null;
    if (EvaluationHelper.equals(parameterType, ParameterType.APPLICATION_ADDRESS_1)) {
      builder =
          ParameterValueApplicationAddress1.staticParseParameterValueBuilder(
              readBuffer, parameterType, numBytes);
    } else if (EvaluationHelper.equals(parameterType, ParameterType.APPLICATION_ADDRESS_2)) {
      builder =
          ParameterValueApplicationAddress2.staticParseParameterValueBuilder(
              readBuffer, parameterType, numBytes);
    } else if (EvaluationHelper.equals(parameterType, ParameterType.INTERFACE_OPTIONS_1)) {
      builder =
          ParameterValueInterfaceOptions1.staticParseParameterValueBuilder(
              readBuffer, parameterType, numBytes);
    } else if (EvaluationHelper.equals(parameterType, ParameterType.BAUD_RATE_SELECTOR)) {
      builder =
          ParameterValueBaudRateSelector.staticParseParameterValueBuilder(
              readBuffer, parameterType, numBytes);
    } else if (EvaluationHelper.equals(parameterType, ParameterType.INTERFACE_OPTIONS_2)) {
      builder =
          ParameterValueInterfaceOptions2.staticParseParameterValueBuilder(
              readBuffer, parameterType, numBytes);
    } else if (EvaluationHelper.equals(
        parameterType, ParameterType.INTERFACE_OPTIONS_1_POWER_UP_SETTINGS)) {
      builder =
          ParameterValueInterfaceOptions1PowerUpSettings.staticParseParameterValueBuilder(
              readBuffer, parameterType, numBytes);
    } else if (EvaluationHelper.equals(parameterType, ParameterType.INTERFACE_OPTIONS_3)) {
      builder =
          ParameterValueInterfaceOptions3.staticParseParameterValueBuilder(
              readBuffer, parameterType, numBytes);
    } else if (EvaluationHelper.equals(parameterType, ParameterType.CUSTOM_MANUFACTURER)) {
      builder =
          ParameterValueCustomManufacturer.staticParseParameterValueBuilder(
              readBuffer, parameterType, numBytes);
    } else if (EvaluationHelper.equals(parameterType, ParameterType.SERIAL_NUMBER)) {
      builder =
          ParameterValueSerialNumber.staticParseParameterValueBuilder(
              readBuffer, parameterType, numBytes);
    } else if (EvaluationHelper.equals(parameterType, ParameterType.CUSTOM_TYPE)) {
      builder =
          ParameterValueCustomTypes.staticParseParameterValueBuilder(
              readBuffer, parameterType, numBytes);
    } else if (true) {
      builder =
          ParameterValueRaw.staticParseParameterValueBuilder(readBuffer, parameterType, numBytes);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "parameterType="
              + parameterType
              + "]");
    }

    readBuffer.closeContext("ParameterValue");
    // Create the instance
    ParameterValue _parameterValue = builder.build(numBytes);

    return _parameterValue;
  }

  public interface ParameterValueBuilder {
    ParameterValue build(Short numBytes);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ParameterValue)) {
      return false;
    }
    ParameterValue that = (ParameterValue) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
